package config

import "os"

var (
	Port    = os.Getenv("app_port")
	AppName = os.Getenv("app_name")
)

var (
	MongoConnectionUri0 = os.Getenv("mongoDbUrl_connection0")
	MongoConntionUri1   = os.Getenv("mongoDbUrl_connection1")
	MongoDbName0        = os.Getenv("dbName_connection0")
	MongoDbName1        = os.Getenv("dbName_connection1")
)

var (
	AWSSecretMongoDB = os.Getenv("awssecrets_mongodb")
)
