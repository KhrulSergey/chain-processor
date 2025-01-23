package service

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/khrulsergey/chain-processor/config"
	"github.com/khrulsergey/chain-processor/logger"
	"github.com/segmentio/kafka-go"
	"reflect"
)

type KafkaService interface {
	ProduceMessage(messageData interface{}) error
}

type kafkaService struct {
	producer    *kafka.Writer
	kafkaConfig *config.KafkaConfig
	log         logger.Logger
}

func NewKafkaService(kafkaConfig *config.KafkaConfig, log logger.Logger) KafkaService {
	producer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaConfig.BootstrapServer),
		Balancer: &kafka.LeastBytes{},
		Transport: &kafka.Transport{
			ClientID: kafkaConfig.ListenerProducerId,
		},
	}
	return &kafkaService{
		kafkaConfig: kafkaConfig,
		producer:    producer,
		log:         log,
	}
}

func (ks *kafkaService) ProduceMessage(messageData interface{}) error {
	dataType := reflect.TypeOf(messageData)
	var topic = ks.kafkaConfig.TopicByTypeMap[dataType]
	if topic == "" {
		return fmt.Errorf("unknown model type: %v for send message to topic", dataType)
	}
	return ks.sendToKafka(messageData, topic)
}

func (ks *kafkaService) sendToKafka(messageData interface{}, topic string) error {
	keyBytes, err := json.Marshal(uuid.New())
	if err != nil {
		ks.log.Errorf("Failed to marshal key: %v", err)
		return err
	}
	encodedData, err := json.Marshal(messageData)
	if err != nil {
		return err
	}

	return ks.producer.WriteMessages(context.Background(),
		kafka.Message{
			Topic: topic,
			Key:   keyBytes,
			Value: encodedData,
		},
	)
}
