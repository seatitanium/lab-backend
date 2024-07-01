package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleRegister(ctx *gin.Context) *errors.CustomErr {
	conn := utils.GetDBConn()

	var object RegisterRequest
	if err := ctx.ShouldBindJSON(&object); err != nil {
		return errors.WrongParam()
	}

	exists, customErr := utils.IsUsernameUsed(object.Username)

	if customErr != nil {
		return customErr
	}

	if exists {
		return errors.DuplicatedUser()
	}

	hash, err := utils.GenerateHash(object.Password)

	if err != nil {
		return errors.ServerError(err)
	}

	result := conn.Create(&utils.Users{
		Username: object.Username,
		Hash:     hash,
		MCID:     object.MCID,
		Email:    object.Email,
	})

	if result.Error != nil {
		return errors.DbError(result.Error)
	}

	middleware.RespSuccess(ctx)
	return nil
}
