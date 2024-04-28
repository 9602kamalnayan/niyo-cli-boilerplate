package cmd

import (
	GLogger "<MODULE_NAME>/lib/logger"
	"<MODULE_NAME>/src/config"
	"<MODULE_NAME>/src/constants"
	awssecretsmanager "<MODULE_NAME>/src/pkg/aws/secretsmanager"
	pkgKafka "<MODULE_NAME>/src/pkg/kafka"
	"context"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const awsRegion = "ap-south-1"

func GetKafkaConfig(ctx context.Context, logger *GLogger.LoggerService) (kafka.ConfigMap, error) {
	var kafkaConfig = kafka.ConfigMap{
		"bootstrap.servers":     config.KafkaBroker,
		"security.protocol":     "SASL_SSL",
		"sasl.mechanisms":       "PLAIN",
		"group.id":              config.KafkaGroupID,
		"heartbeat.interval.ms": 10000,
	}
	if config.AppEnv == constants.EnvTesting {
		kafkaConfig["sasl.username"] = config.KafkaClsaKey
		kafkaConfig["sasl.password"] = config.KafkaClsaSecret
	} else {
		smClient := awssecretsmanager.GetSecretManagerClient(ctx, awsRegion, logger)
		kafkaSecrets, err := smClient.GetSecret(ctx, config.KafkaSecretArn)
		if err != nil {
			logger.Error(ctx, "error_occured_while_fetching_kafka_secrets", err)
			return nil, err
		}
		kafkaConfig["sasl.username"] = kafkaSecrets["clsa_key"]
		kafkaConfig["sasl.password"] = kafkaSecrets["clsa_secret"]
	}
	return kafkaConfig, nil
}

func InitializeKafkaProducer(ctx context.Context, logger *GLogger.LoggerService) error {
	var kafkaConfig, err = GetKafkaConfig(ctx, logger)
	if err != nil {
		logger.Error(ctx, "error_while_creating_kafka_config", err)
		return err
	}
	_, err = pkgKafka.InitializeKafkaProducer(ctx, &kafkaConfig, logger)
	if err != nil {
		logger.Error(ctx, "error_occurred_while_initializing_kafka_producer", err)
		return err
	}
	return nil
}

func InitializeKafkaConsumer(ctx context.Context, consumerName string, kafkaTopics []string, uniqueMessageHandler pkgKafka.UniqueKafkaMessageHandler, logger *GLogger.LoggerService) error {
	var kafkaConfig, err = GetKafkaConfig(ctx, logger)
	if err != nil {
		logger.Error(ctx, "error_while_creating_kafka_config", err)
		return err
	}
	consumer := pkgKafka.InitializeKafkaConsumer(fmt.Sprintf("%s:%s:%s", config.AppName, constants.AppNameSuffixConsumer, consumerName), kafkaTopics, uniqueMessageHandler, logger)
	err = consumer.StartConsumer(ctx, &kafkaConfig)
	if err != nil {
		logger.Error(ctx, "error_while_running_kafka_consumer", err)
		return err
	}
	return nil
}
