package main

import (
	"$ServiceName/src/cmd"
	"$ServiceName/src/config"
	"$ServiceName/src/constants"
	appgin "$ServiceName/src/internal/app/gin"
	"$ServiceName/src/internal/helpers"
	"$ServiceName/src/modules/database"
	"context"
	"fmt"
	"log"

	"github.com/pkg/errors"
)

const defaultErrorMessage = "could_not_start_gin_server"

func main() {
	logger := helpers.InitLogger()
	ctx := context.Background()
	ctx = context.WithValue(ctx, config.AppName, fmt.Sprintf("%s:%s", config.AppName, constants.AppNameSuffixHTTPServer))
	db, err := cmd.EstablishDBConnection(ctx, logger)
	defer func(db database.DBInterface) {
		_ = db.Disconnect(ctx)
	}(db)
	server := appgin.SetupServer(ctx, logger)
	err = server.Run(config.Port)
	if err != nil {
		err = errors.Wrap(err, "error_occurred_while_starting_gin_server")
		logger.Error(ctx, err)
		log.Fatal(defaultErrorMessage + err.Error())
	}
}
