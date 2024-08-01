package middleware

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func SelfCheck(ctx *gin.Context) {
	if !utils.IsSelfOnly(ctx.Request.RequestURI) {
		return
	}

	parsed, parseErr := utils.GetPayloadFromToken(ctx.Request.Header.Get("Authorization"))

	if parseErr != nil {
		handlers.RespForbidden(ctx)
		ctx.Abort()
		return
	}

	username := ctx.DefaultQuery("username", "")
	playername := ctx.DefaultQuery("playername", "")

	// Playername can only be used by admin in self-only endpoints.
	if playername != "" && !utils.IsAdmin(parsed.Username) {
		handlers.RespForbidden(ctx)
		ctx.Abort()
		return
	}

	// Check if the requested username is identical to that in token.
	if parsed.Username != username && !utils.IsAdmin(parsed.Username) {
		handlers.RespForbidden(ctx)
		ctx.Abort()
		return
	}
}
