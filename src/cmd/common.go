package cmd

import (
	GLogger "<MODULE_NAME>/lib/logger"
	"<MODULE_NAME>/src/config"
	"<MODULE_NAME>/src/constants"
	"<MODULE_NAME>/src/modules/database"
	"context"
	"github.com/pkg/errors"
	"log"
)

func EstablishDBConnection(ctx context.Context, logger *GLogger.LoggerService) (database.DBInterface, error) {
	var db database.DBInterface
	if config.AppEnv == constants.EnvTesting || config.AppEnv == constants.EnvDevelopment || config.AppEnv == constants.EnvStaging {
		db = database.NewDBInstance(logger, config.Connections.Connection0.ConnectionUri, database.OpenURI, "", database.ConnectionOptions{}, "")
	} else {
		db = database.NewDBInstance(logger, config.Connections.Connection0.ConnectionUri, database.AwsSecrets, config.AWSSecretMongoDB, database.ConnectionOptions{}, config.Connections.Connection0.DbName)
	}
	err := db.Connect(ctx)
	if err != nil {
		err = errors.Wrap(err, "error_encountered_with_db_connection")
		logger.Error(ctx, err)
		log.Fatal(err)
	}
	return db, err
}
