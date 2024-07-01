package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleCheck(ctx *gin.Context) *errors.CustomErr {
	token := ctx.Query("token")

	if token == "" {
		return errors.UnAuth()
	}

	err := utils.CheckJWT(token)

	if err != nil {
		return err
	}

	middleware.RespSuccess(ctx)
	return nil
}
