package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
	"time"
)

func HandleLogin(ctx *gin.Context) *errHandler.CustomErr {
	var object LoginRequest
	if err := ctx.ShouldBindJSON(&object); err != nil {
		return errHandler.WrongParam()
	}

	user, customErr := utils.GetUserByUsername(object.Username)
	if customErr != nil {
		return customErr
	}

	if utils.VerifyHash([]byte(user.Hash), []byte(object.Password)) {
		jwt, err := utils.GenerateJWT(utils.JWTPayload{
			Username:  object.Username,
			UpdatedAt: time.Now().UnixMilli(),
		})

		if err != nil {
			return err
		}

		middleware.RespSuccess(ctx, jwt)
	} else {
		return errHandler.UnAuth()
	}

	return nil
}
