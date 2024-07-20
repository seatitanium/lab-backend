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

	// Header naming conventions considered from https://stackoverflow.com/questions/3561381/custom-http-headers-naming-conventions
	if ctx.Request.Header.Get("Seati-Server-Secret") != utils.GlobalConfig.ServerSecret {
		handlers.RespForbidden(ctx)
		ctx.Abort()
		return
	}
}
