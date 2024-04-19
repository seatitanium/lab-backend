package middleware

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/utils"
)

func TokenCheck(ctx *gin.Context) {
	if !utils.NeedAuthorize(ctx.HandlerName()) {
		return
	}

	checkErr := utils.CheckJWT(ctx.Request.Header.Get("Token"))

	RespTokenError(ctx, checkErr.Code, checkErr.Msg)
}
