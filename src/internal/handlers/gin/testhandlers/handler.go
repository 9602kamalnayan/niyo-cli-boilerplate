package testhandlers

import (
	"$ServiceName/src/internal/web"
	"github.com/gin-gonic/gin"
)

func CheckRouterPath(ctx *gin.Context) (*web.JSONResponse, *web.ErrorStruct) {
	return web.NewHTTPSuccessResponse(ctx, "successfully called CheckRouterPath", map[string]interface{}{}), nil
}
