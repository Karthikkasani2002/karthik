package config

import "os"

type Config struct {

	KafkaBrokers string
	KafkaTopic string
	KafkaCA string
	KafkaUsername string
        KafkaPassword string
	PostgresHost string
	PostgresPort string
	PostgresDB string
	PostgresUser string
	PostgresPass string
	PostgresCA string

	RedisHost string
	RedisPass string
	RedisCA string
	RedisPort string

	ListenAddr string

	ServiceName string
	LogLevel string
}

func must(key string) string {

	v := os.Getenv(key)

	if v == "" {

		panic("missing env " + key)
	}

	return v
}

func Load() Config {

	return Config{

		KafkaBrokers: must("KAFKA_BROKERS"),

		KafkaTopic: must("TOPIC_ONBOARD_INIT"),

		KafkaCA: must("KAFKA_CA"),

		KafkaUsername: must("KAFKA_USERNAME"),

                KafkaPassword: must("KAFKA_PASSWORD"),

		PostgresHost: must("POSTGRES_HOST"),

		PostgresPort: must("POSTGRES_PORT"),

		PostgresDB: must("POSTGRES_DB"),

		PostgresUser: must("POSTGRES_USER"),

		PostgresPass: must("POSTGRES_PASSWORD"),

		PostgresCA: must("POSTGRES_CA"),

		RedisHost: must("REDIS_HOST"),

		RedisPass: must("REDIS_PASSWORD"),

		RedisPort: must("REDIS_PORT"),

		RedisCA: must("REDIS_CA"),

		ListenAddr: must("LISTEN_ADDR"),

		ServiceName: "onboarding-initiator",

		LogLevel: "info",
	}
}
