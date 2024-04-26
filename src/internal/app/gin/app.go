package app_gin

import (
	GLogger "<MODULE_NAME>/lib/logger"
	"<MODULE_NAME>/src/config"
	apigin "<MODULE_NAME>/src/internal/api/gin"
	"context"
	"github.com/gin-gonic/gin"
)

func SetupServer(ctx context.Context, log *GLogger.LoggerService) *gin.Engine {
	log.Info(ctx, "setting_up_routes...")
	gin.SetMode(config.GinMode)
	engine := gin.New()
	engine.HandleMethodNotAllowed = true
	// engine.Use(ginmiddleware.OtelMiddleware(config.ServiceName))
	engine.Use(gin.Recovery())
	//Create Your Route here
	baseRouterGroup := &apigin.RouterGroup{Router: engine.Group(config.BaseRouterPath), Log: log}
	baseRouterGroup.DefaultRoutes()
	baseRouterGroup.AddV1Routes()
	return engine
}
