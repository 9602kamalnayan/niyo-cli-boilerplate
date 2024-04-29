//go:build uat
// +build uat

package config

import (
	GLogger "$ServiceName/lib/logger"
	"$ServiceName/src/constants"
	"fmt"

	"github.com/gin-gonic/gin"
)

const LogLevel = GLogger.DEBUG
const GinMode = gin.DebugMode
const AppEnv = constants.EnvStaging

var BaseRouterPath = fmt.Sprintf("/%s/%s/%s", "SET_YOUR_BASE_PREFIX", AppEnv, "SET_YOUR_BASE_PREFIX")

type ConnectionParamsStruct struct {
	ConnectionUri string
	DbName        string
}
type ConnectionsStruct struct {
	Connection0 ConnectionParamsStruct
	Connection1 ConnectionParamsStruct
}

var Connections = ConnectionsStruct{
	Connection0: ConnectionParamsStruct{
		ConnectionUri: MongoConnectionUri0,
		DbName:        MongoDbName0,
	},
	Connection1: ConnectionParamsStruct{
		ConnectionUri: MongoConntionUri1,
		DbName:        MongoDbName1,
	},
}
