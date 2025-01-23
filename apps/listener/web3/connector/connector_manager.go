package connector

import (
	"fmt"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"sync"
	"time"
)

type ChainConnectorManager struct {
	chainConfigs *config.ChainConfigs
	logger       logger.Logger
	connectors   map[int]ChainConnector
	executors    sync.Map
	mu           sync.Mutex
}

func NewChainConnectorManager(chainConfigs *config.ChainConfigs, logger logger.Logger) *ChainConnectorManager {
	manager := &ChainConnectorManager{
		chainConfigs: chainConfigs,
		logger:       logger,
		connectors:   make(map[int]ChainConnector),
	}
	manager.ConnectToChains()
	return manager
}

func (m *ChainConnectorManager) ConnectToChains() {
	for _, chainConfig := range m.chainConfigs.GetAllEnabledChains() {
		if !chainConfig.Enabled {
			continue
		}
		conn := NewEVMConnector(chainConfig, m.logger)
		err := conn.Connect()
		if err != nil {
			m.logger.Errorf("Failed to connect to chain id %d: %v", chainConfig.ChainId, err)
		}
		m.mu.Lock()
		m.connectors[chainConfig.ChainId] = conn
		m.mu.Unlock()

		// Run accessibility check in a separate goroutine
		go m.initAccessibilityScheduler(chainConfig.ChainId)
	}
}

func (m *ChainConnectorManager) GetChainConnectorSize() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return len(m.connectors)
}

func (m *ChainConnectorManager) GetChainConnector(chainId int) (ChainConnector, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	conn, exists := m.connectors[chainId]
	if !exists {
		return nil, fmt.Errorf("connector for chain id %d not found", chainId)
	}
	return conn, nil
}

func (m *ChainConnectorManager) initAccessibilityScheduler(chainId int) {
	chainConfig, err := m.chainConfigs.GetConfigByChainId(chainId)
	if err != nil {
		m.logger.Errorf("Chain config not found for chain id %d. Original error: %v", chainId, err)
		return
	}

	ticker := time.NewTicker(time.Duration(chainConfig.AccessibilityTimeoutSec) * time.Second)
	m.executors.Store(chainId, ticker)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			conn, err := m.GetChainConnector(chainId)
			if err != nil {
				m.logger.Errorf("Error getting connector: %v", err)
				return
			}
			if err := conn.CheckAccessibility(); err != nil {
				m.logger.Errorf("Accessibility check failed for chain id %d: %v", chainId, err)
				m.exceptionNotificationHandler(chainId, err)
			}
		}
	}
}

func (m *ChainConnectorManager) exceptionNotificationHandler(chainId int, connectorException error) {
	chainConfig, err := m.chainConfigs.GetConfigByChainId(chainId)
	if err != nil {
		m.logger.Errorf("Chain config not found for chain id %d. Original error: %v", chainId, err)
		return
	}

	if connectorException != nil {
		restoreTimeoutSec := chainConfig.RestoreTimeoutSec
		m.logger.Infof("Try to restore connection to chain: %d within %d seconds.", chainId, restoreTimeoutSec)

		m.disposeAccessibilityScheduler(chainId)
		time.Sleep(time.Duration(restoreTimeoutSec) * time.Second)

		conn, err := m.GetChainConnector(chainId)
		if err != nil {
			m.logger.Errorf("Error getting connector: %v", err)
			return
		}
		conn.Dispose()
		err = conn.Connect()
		if err == nil {
			// Handle system events (pseudo-code, replace with actual implementation)
			// systemEventsStream.handleEvent(...)
		} else {
			m.logger.Warningf("New unexpected error in connector for chain: %d, exception: %v", chainId, err)
			return
		}
	} else {
		m.logger.Warningf("New unexpected error in connector for chain: %d, exception: %v", chainId, connectorException)
		return
	}

	go m.initAccessibilityScheduler(chainId)
}

func (m *ChainConnectorManager) disposeAccessibilityScheduler(chainId int) {
	if ticker, ok := m.executors.Load(chainId); ok {
		ticker.(*time.Ticker).Stop()
		m.executors.Delete(chainId)
	}
}
