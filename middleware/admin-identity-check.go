package middleware

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"slices"
)

func AdminCheck(ctx *gin.Context) {
	if !utils.IsAdminOnly(ctx.Request.RequestURI) {
		return
	}

	parsed, parseErr := utils.GetPayloadFromToken(ctx.Request.Header.Get("Authorization"))

	if parseErr != nil {
		handlers.RespForbidden(ctx)
		ctx.Abort()
		return
	}

	if !slices.Contains(utils.GlobalConfig.Administrators, parsed.Username) {
		handlers.RespForbidden(ctx)
		ctx.Abort()
		return
	}
}
