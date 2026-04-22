package kafka

import (

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"

	"onboarding/internal/config"
)

type Producer struct {

	client *kafka.Producer

	topic string
}

func New(cfg config.Config) *Producer {

	p, err := kafka.NewProducer(&kafka.ConfigMap{

		"bootstrap.servers": cfg.KafkaBrokers,

		"security.protocol": "SASL_SSL",

		"ssl.ca.location": cfg.KafkaCA,

		"ssl.endpoint.identification.algorithm": "none",

		"sasl.mechanism": "SCRAM-SHA-512",

                "sasl.username": cfg.KafkaUsername,
		
                "sasl.password": cfg.KafkaPassword,

		"enable.idempotence": true,

		"acks": "all",

		"compression.type": "snappy",
	})

	if err != nil {

		panic(err)
	}

	return &Producer{

		client: p,

		topic: cfg.KafkaTopic,
	}
}

func (p *Producer) Publish(key string, value []byte) error {

	return p.client.Produce(&kafka.Message{

		TopicPartition: kafka.TopicPartition{

			Topic: &p.topic,

			Partition: kafka.PartitionAny,
		},

		Key: []byte(key),

		Value: value,

	}, nil)
}

func (p *Producer) Close() {

	p.client.Flush(5000)

	p.client.Close()
}
