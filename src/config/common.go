package config

import "os"

var (
	Port    = os.Getenv("app_port")
	AppName = os.Getenv("app_name")
)

var (
	KafkaBroker     = os.Getenv("kafka_broker")
	KafkaGroupID    = os.Getenv("kafka_groupid")
	KafkaSecretArn  = os.Getenv("kafka_secretarn")
	KafkaClsaKey    = os.Getenv("kafka_clsakey")
	KafkaClsaSecret = os.Getenv("kafka_clsasecret")
)
