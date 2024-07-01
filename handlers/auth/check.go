package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
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

	handlers.RespSuccess(ctx)
	return nil
}
