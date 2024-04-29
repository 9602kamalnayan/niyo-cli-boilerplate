//go:build beta
// +build beta

package config

import (
	GLogger "$ServiceName/lib/logger"
	"$ServiceName/src/constants"
	"fmt"

	"github.com/gin-gonic/gin"
)

const LogLevel = GLogger.DEBUG
const GinMode = gin.DebugMode
const AppEnv = constants.EnvBeta

var BaseRouterPath = fmt.Sprintf("/%s/%s/%s", "SET_YOUR_BASE_PREFIX", AppEnv, "SET_YOUR_BASE_PREFIX")
