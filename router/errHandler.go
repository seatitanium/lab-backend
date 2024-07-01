package router

import (
	"fmt"
	"seatimc/backend/errors"
	"seatimc/backend/utils"

	"github.com/gin-gonic/gin"
)

type HandlerFunc func(c *gin.Context) *errors.CustomErr

func requestInfo(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.String())
}

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		customErr := handler(c)
		if customErr != nil {
			var exception *errors.ApiException
			exception = customErr.Handle()
			exception.Request = requestInfo(c)
			c.JSON(exception.HttpCode, exception)
		}
	}
}

func handleNotFound(ctx *gin.Context) {
	handleErr := errors.NotFound()
	handleErr.Request = requestInfo(ctx)
	ctx.JSON(handleErr.HttpCode, handleErr)
}

func handleVersion(ctx *gin.Context) {
	ctx.String(200, "tisea @ "+utils.GlobalConfig.Version)
}

func handleUnauth(ctx *gin.Context) {
	handleErr := errors.UnAuth()
	handleErr.Request = requestInfo(ctx)
	ctx.JSON(handleErr.HttpCode, handleErr)
}
