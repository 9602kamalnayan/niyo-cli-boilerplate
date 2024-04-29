//go:build prod
// +build prod

package config

import (
	GLogger "$ServiceName/lib/logger"
	"$ServiceName/src/constants"
	"fmt"

	"github.com/gin-gonic/gin"
)

const LogLevel = GLogger.DEBUG
const GinMode = gin.DebugMode
const AppEnv = constants.EnvProd

var BaseRouterPath = fmt.Sprintf("/%s/%s/%s", "SET_YOUR_BASE_PREFIX", AppEnv, "SET_YOUR_BASE_PREFIX")
