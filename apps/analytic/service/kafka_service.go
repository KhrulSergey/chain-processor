package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/khrulsergey/chain-processor/pkg/dto"
	"github.com/segmentio/kafka-go"
	"reflect"
)

type KafkaService interface {
	listenMessages(topic string) error
}

type kafkaService struct {
	kafkaConfig   *config.KafkaConfig
	consumersMap  map[string]*kafka.Reader
	dataProcessor DataProcessor
	log           logger.Logger
}

func NewKafkaService(kafkaConfig *config.KafkaConfig, dataProcessor DataProcessor, log logger.Logger) KafkaService {
	consumersMap := make(map[string]*kafka.Reader)
	for topic, _ := range kafkaConfig.TypeByTopicMap {
		consumersMap[topic] = initKafkaReader(topic, kafkaConfig)
	}

	ks := &kafkaService{
		kafkaConfig:   kafkaConfig,
		consumersMap:  consumersMap,
		dataProcessor: dataProcessor,
		log:           log,
	}

	for topic, _ := range kafkaConfig.TypeByTopicMap {
		go func(topic string) {
			err := ks.listenMessages(topic)
			if err != nil {
				ks.log.Errorf("failed to connect to topic: %s, error: %v", topic, err)
			}
		}(topic)
	}
	return ks
}

func initKafkaReader(topic string, kafkaConfig *config.KafkaConfig) *kafka.Reader {
	readerConfig := kafka.ReaderConfig{
		Brokers:     []string{kafkaConfig.BootstrapServer},
		GroupID:     kafkaConfig.AnalyticConsumerId,
		Topic:       topic,
		MaxAttempts: 4,
	}
	return kafka.NewReader(readerConfig)
}

func (ks kafkaService) listenMessages(topic string) error {
	consumer := ks.consumersMap[topic]
	if consumer == nil {
		err := fmt.Errorf("no consumer for topic: %v", topic)
		ks.log.Fatalf(err.Error())
		return err
	}
	ks.log.Infof("start listen messages on topic: %s", topic)
	ctx := context.Background()
	for {
		m, err := consumer.FetchMessage(ctx)
		if err != nil {
			return err
		}
		ks.log.Debugf("message at topic/partition/offset %v/%v/%v: %s = %s\n",
			m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
		dataType := ks.kafkaConfig.TypeByTopicMap[topic]
		data := reflect.New(dataType.Elem()).Interface()
		err = json.Unmarshal(m.Value, data)
		if err != nil {
			return err
		}
		var externalId string
		err = json.Unmarshal(m.Key, &externalId)
		if err != nil {
			return err
		}
		data.(dto.BaseDto).SetExternalId(externalId)
		err = ks.dataProcessor.Handle(data)
		if err != nil {
			if err = consumer.CommitMessages(ctx, m); err != nil {
				ks.log.Warning("failed to commit message:", err)
			}
		}
	}
}
