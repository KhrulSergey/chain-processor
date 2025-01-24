package service

import (
	"context"
	"errors"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/segmentio/kafka-go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"reflect"
	"testing"
)

// MockKafkaWriter mocks the kafkaWriter interface
type MockKafkaWriter struct {
	mock.Mock
}

func (m *MockKafkaWriter) WriteMessages(ctx context.Context, msgs ...kafka.Message) error {
	args := m.Called(ctx, msgs)
	return args.Error(0)
}

func TestProduceMessageSuccess(t *testing.T) {
	// Setup
	mockKafkaWriter := new(MockKafkaWriter)
	mockLogger := new(logger.MockLogger)

	kafkaConfig := &config.KafkaConfig{
		BootstrapServer:    "localhost:9092",
		ListenerProducerId: "test-producer",
		TopicByTypeMap: map[reflect.Type]string{
			reflect.TypeOf("test message"): "test-topic",
		},
	}

	message := "test message"
	topic := "test-topic" // Ensure 'topic' is declared here

	// Mock KafkaWriter behavior
	mockKafkaWriter.
		On("WriteMessages", mock.Anything,
			mock.MatchedBy(func(msgs []kafka.Message) bool {
				return len(msgs) == 1 && msgs[0].Topic == topic
			})).
		Return(nil)

	kafkaService := &kafkaService{
		producer:    mockKafkaWriter,
		kafkaConfig: kafkaConfig,
		log:         mockLogger,
	}

	// Act
	err := kafkaService.ProduceMessage(message)

	// Assert
	assert.NoError(t, err)
	mockKafkaWriter.AssertExpectations(t)
}

func TestProduceMessageUnknownType(t *testing.T) {
	// Setup
	mockKafkaWriter := new(MockKafkaWriter)
	mockLogger := new(logger.MockLogger)

	kafkaConfig := &config.KafkaConfig{
		BootstrapServer:    "localhost:9092",
		ListenerProducerId: "test-producer",
		TopicByTypeMap:     map[reflect.Type]string{},
	}

	message := 12345 // Unknown type

	kafkaService := &kafkaService{
		producer:    mockKafkaWriter,
		kafkaConfig: kafkaConfig,
		log:         mockLogger,
	}

	// Act
	err := kafkaService.ProduceMessage(message)

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unknown model type")
}

func TestSendToKafkaFailure(t *testing.T) {
	// Setup
	mockKafkaWriter := new(MockKafkaWriter)
	mockLogger := new(logger.MockLogger)

	kafkaConfig := &config.KafkaConfig{
		BootstrapServer:    "localhost:9092",
		ListenerProducerId: "test-producer",
		TopicByTypeMap: map[reflect.Type]string{
			reflect.TypeOf("test message"): "test-topic",
		},
	}

	message := "test message"

	// Mock KafkaWriter behavior
	mockKafkaWriter.
		On("WriteMessages", mock.Anything, mock.Anything).
		Return(errors.New("kafka write error"))

	// Expect Errorf to be called with the error message
	mockLogger.
		On("Errorf", mock.Anything, mock.Anything).
		Return()

	kafkaService := &kafkaService{
		producer:    mockKafkaWriter,
		kafkaConfig: kafkaConfig,
		log:         mockLogger,
	}

	// Act
	err := kafkaService.ProduceMessage(message)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, "kafka write error", err.Error())
	mockKafkaWriter.AssertExpectations(t)
	mockLogger.AssertExpectations(t)
}

func TestSendToKafkaSerializationError(t *testing.T) {
	// Setup
	mockKafkaWriter := new(MockKafkaWriter)
	mockLogger := new(logger.MockLogger)

	kafkaConfig := &config.KafkaConfig{
		BootstrapServer:    "localhost:9092",
		ListenerProducerId: "test-producer",
		TopicByTypeMap: map[reflect.Type]string{
			reflect.TypeOf(make(chan int)): "test-topic", // Define Invalid type
		},
	}

	message := make(chan int) // Cannot be serialized by json.Marshal

	// Expect Errorf to be called with the serialization error message
	mockLogger.
		On("Errorf", mock.Anything, mock.Anything).
		Return()

	kafkaService := &kafkaService{
		producer:    mockKafkaWriter,
		kafkaConfig: kafkaConfig,
		log:         mockLogger,
	}

	// Act
	err := kafkaService.ProduceMessage(message)

	// Assert
	assert.Error(t, err)
	mockKafkaWriter.AssertNotCalled(t, "WriteMessages", mock.Anything, mock.Anything)
	mockLogger.AssertExpectations(t)
}
