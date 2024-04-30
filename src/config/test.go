//go:build !dev && !uat && !beta && !prod
// +build !dev,!uat,!beta,!prod

package config

import (
	GLogger "$ServiceName/lib/logger"
	"$ServiceName/src/constants"
	"fmt"

	"github.com/gin-gonic/gin"
)

const LogLevel = GLogger.DEBUG
const GinMode = gin.DebugMode
const AppEnv = constants.EnvTesting

var BaseRouterPath = fmt.Sprintf("/%s/%s/%s", "SET_YOUR_BASE_PREFIX", AppEnv, "SET_YOUR_BASE_PREFIX")
