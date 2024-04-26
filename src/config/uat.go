package config

import (
	GLogger "<MODULE_NAME>/lib/logger"
	"<MODULE_NAME>/src/constants"
	"fmt"

	"github.com/gin-gonic/gin"
)

const LogLevel = GLogger.DEBUG
const GinMode = gin.DebugMode
const AppEnv = constants.EnvStaging

var BaseRouterPath = fmt.Sprintf("/%s/%s/%s", "SET_YOUR_BASE_PREFIX", AppEnv, "SET_YOUR_BASE_PREFIX")
