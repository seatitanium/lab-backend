package middleware

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func ServerCheck(ctx *gin.Context) {
	if !utils.IsServerOnly(ctx.Request.RequestURI) {
		return
	}

	if !utils.VerifyServerSecretCtx(ctx) {
		handlers.RespForbidden(ctx)
		ctx.Abort()
		return
	}
}
