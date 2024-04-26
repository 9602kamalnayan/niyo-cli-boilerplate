package api_gin

import (
	GLogger "$SERVICE_NAME/lib/logger"
	"$SERVICE_NAME/src/constants"
	routemodels "$SERVICE_NAME/src/internal/api/gin/models"
	"$SERVICE_NAME/src/internal/handlers/gin/testhandlers"
	ginmiddlewares "$SERVICE_NAME/src/internal/middleware/gin"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
)

type RouterGroup struct {
	Router *gin.RouterGroup
	Log    *GLogger.LoggerService
}

func (rg *RouterGroup) DefaultRoutes() {
	defaultRoutes := rg.Router.Group(string(constants.RGDocs))
	defaultRoutes.GET("/", ginSwagger.WrapHandler(swaggerFiles.Handler))
	defaultRoutes.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, map[string]interface{}{
			"message": "successfully_called_health_endpoint",
		})
		ctx.Next()
	})
}
func (rg *RouterGroup) AddV1Routes() {
	v1Routes := rg.Router.Group(string(constants.RV1))
	internalRoutes := v1Routes.Group(string(constants.RGInternalGroup))
	internalRoutes.GET(string(constants.RGCheckRouterPath), ginmiddlewares.ServeEndpoint[routemodels.TestModel](testhandlers.CheckRouterPath))

}
