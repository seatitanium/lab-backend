package middleware

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func TokenCheck(ctx *gin.Context) {
	if !utils.NeedAuthorize(ctx.Request.RequestURI) {
		return
	}

	checkErr := utils.CheckJWT(ctx.Request.Header.Get("Authorization"))

	if checkErr != nil {
		handlers.RespTokenError(ctx, checkErr.Code, checkErr.Msg)
		ctx.Abort()
		return
	}
}
