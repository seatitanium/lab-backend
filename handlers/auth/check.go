package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleCheck(ctx *gin.Context) *errHandler.CustomErr {
	token := ctx.Query("token")

	if token == "" {
		return errHandler.UnAuth()
	}

	err := utils.CheckJWT(token)

	if err != nil {
		return err
	}

	middleware.RespSuccess(ctx)
	return nil
}
