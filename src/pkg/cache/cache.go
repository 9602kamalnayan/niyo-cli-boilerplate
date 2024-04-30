package cache

import (
	"$ServiceName/src/config"
	"$ServiceName/src/constants"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
)

var defaultErrorMessage = "UNEXPECTED_REDIS_CLIENT_TYPE"
var ctx = context.Background()
var RedisClient interface{} = nil

func getRedisClient(log *logrus.Entry) error {
	if RedisClient == nil {
		redisClient, err := InitializeRedisClient(log)
		if err != nil {
			log.Error("ERROR_WHILE_ESTABLISH_REDIS_SERVER ", err)
			return err
		}
		RedisClient = redisClient
	}
	return nil
}
func InitializeRedisClient(log *logrus.Entry) (interface{}, error) {
	log.Info("ATTEMPTING_TO_ESTABLISH_REDIS_CONNECTION")
	if config.AppEnv != constants.EnvTesting {
		redisClient := redis.NewClusterClient(&redis.ClusterOptions{
			Addrs: []string{config.RedisHost},
		})
		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			log.Error("FAILED_TO_ESTABLISH_REDIS_CONNECTION ", err)
			return nil, err
		}
		RedisClient = redisClient
	} else {
		redisClient := redis.NewClient(&redis.Options{
			Addr: config.RedisHost,
		})
		_, err := redisClient.Ping(ctx).Result()
		if err != nil {
			log.Error("FAILED_TO_ESTABLISH_REDIS_CONNECTION ", err)
			return nil, err
		}
		RedisClient = redisClient
	}
	log.Info("SUCCESSFULLY_ESTABLISHED_REDIS_CONNECTION")
	return RedisClient, nil
}

func SetToCache(key string, val interface{}, log *logrus.Entry) error {
	log.Info(fmt.Sprintf("SETTING_KEY_VALUE_IN_REDIS_CACHE key: %s val: %s ", key, val))
	var res *redis.StatusCmd
	err := getRedisClient(log)
	if err != nil {
		return err
	}
	if rcClient, ok := RedisClient.(*redis.ClusterClient); ok {
		res = rcClient.Set(ctx, key, transformInputToCacheValue(val, log), 0)
	} else if rlClient, ok := RedisClient.(*redis.Client); ok {
		res = rlClient.Set(ctx, key, transformInputToCacheValue(val, log), 0)
	} else {
		err := errors.New(defaultErrorMessage)
		log.Error(defaultErrorMessage, err)
		return err
	}
	if res.Err() != nil {
		log.Error("ERROR_WHILE_SETTING_VALUE_IN_REDIS_CACHE ", res.Err())
		return res.Err()
	}
	log.Debug(fmt.Sprintf("SUCCESSFULLY_SET_KEY_VALUE_IN_REDIS_CACHE %s", key), res.Val())
	return nil
}

func GetFromCache(key string, log *logrus.Entry) string {
	log.Debug("GETTING_FROM_REDIS_CACHE ", key)
	var value string
	var err error
	err = getRedisClient(log)
	if err != nil {
		return value
	}
	if rcClient, ok := RedisClient.(*redis.ClusterClient); ok {
		value, err = rcClient.Get(ctx, key).Result()
	} else if rlClient, ok := RedisClient.(*redis.Client); ok {
		value, err = rlClient.Get(ctx, key).Result()
	} else {
		log.Error(defaultErrorMessage, err)
		return value
	}
	if err != nil {
		if err == redis.Nil {
			log.Debug(fmt.Sprintf("NO_VALUE_FOUND_IN_REDIS_CACHE %s: %s", key, value))
			return value
		} else {
			log.Error("ERROR_WHILE_GETTING_VALUE_FROM_REDIS_CACHE ", err)
		}

		return value
	}
	log.Debug(fmt.Sprintf("SUCCESSFULLY_RECEIVED_VALUE_FROM_REDIS_CACHE %s: %s", key, value))
	return value
}

func DelFromCache(key string, log *logrus.Entry) error {
	log.Debug("DELETING_KEY_FROM_REDIS_CACHE", key)
	err := getRedisClient(log)
	if err != nil {
		return err
	}
	if rcClient, ok := RedisClient.(*redis.ClusterClient); ok {
		_, err := rcClient.Del(ctx, key).Result()
		if err != nil {
			log.Error("ERROR_WHILE_DELETING_VALUE_FROM_REDIS_CACHE ", err)
			return err
		}
	} else if rlClient, ok := RedisClient.(*redis.Client); ok {
		_, err := rlClient.Del(ctx, key).Result()
		if err != nil {
			log.Error("ERROR_WHILE_DELETING_VALUE_FROM_REDIS_CACHE ", err)
			return err
		}
	} else {
		err := errors.New(defaultErrorMessage)
		log.Error(defaultErrorMessage, err)
		return err
	}

	log.Debug("SUCCESSFULLY_DELETED_KEY_FROM_REDIS_CACHE")
	return nil
}
func transformInputToCacheValue(input interface{}, log *logrus.Entry) []byte {
	jsonByte, err := json.Marshal(input)

	if err != nil {
		log.Error("ERROR_WHILE_MARSHALING_FOR_CACHE ", err)
		return nil
	}
	return jsonByte
}
func transformCacheValueToInput(cacheValue string, log *logrus.Entry) []byte {
	return nil
}
