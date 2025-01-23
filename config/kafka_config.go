package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/khrulsergey/chain-processor/pkg/dto"
	"github.com/khrulsergey/chain-processor/pkg/types"
	"reflect"
)

// KafkaConfig contains needed envs to run kafka service
type KafkaConfig struct {
	BootstrapServer      string                  `env:"DESTINATION_KAFKA_BOOTSTRAP_SERVERS" envDefault:"localhost:9092"`
	DeliveryGuarantee    types.DeliveryGuarantee `env:"DESTINATION_KAFKA_DELIVERY_GUARANTEE" envDefault:"AT_MOST_ONCE"`
	ListenerProducerId   string                  `env:"DESTINATION_KAFKA_PRODUCER_ID" envDefault:"chain-listener-service"`
	AnalyticConsumerId   string                  `env:"DESTINATION_KAFKA_CONSUMER_ID" envDefault:"analytics-service"`
	TradeEventTopicName  string                  `env:"DESTINATION_KAFKA_TRADE_EVENT_TOPIC_NAME" envDefault:"chain-trade-event"`
	TokenPairTopicName   string                  `env:"DESTINATION_KAFKA_TOKEN_PAIR_TOPIC_NAME" envDefault:"chain-token-pair"`
	TransactionTopicName string                  `env:"DESTINATION_KAFKA_TRANSACTIONS_TOPIC_NAME" envDefault:"chain-transaction"`
	TypeByTopicMap       map[string]reflect.Type
	TopicByTypeMap       map[reflect.Type]string
}

// NewKafkaConfig parses envs and constructs the kafka config
func NewKafkaConfig() (*KafkaConfig, error) {
	config := &KafkaConfig{}
	if err := env.Parse(config); err != nil {
		return nil, err
	}
	typeByTopicMap := map[string]reflect.Type{
		config.TradeEventTopicName: reflect.TypeOf(&dto.TradeDto{}),
		config.TokenPairTopicName:  reflect.TypeOf(&dto.TokenPairDto{}),
	}
	topicByTypeMap := make(map[reflect.Type]string)
	for topic, dataType := range typeByTopicMap {
		topicByTypeMap[dataType] = topic
	}
	config.TypeByTopicMap = typeByTopicMap
	config.TopicByTypeMap = topicByTypeMap
	return config, nil
}
