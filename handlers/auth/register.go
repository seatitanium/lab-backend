package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleRegister(ctx *gin.Context) *errHandler.CustomErr {
	conn := utils.GetDBConn()

	var object RegisterRequest
	if err := ctx.ShouldBindJSON(&object); err != nil {
		return errHandler.WrongParam()
	}

	hash, err := utils.GenerateHash(object.Password)
	if err != nil {
		return errHandler.ServerError(err)
	}

	result := conn.Create(&utils.Users{
		Username: object.Username,
		Hash:     hash,
		MCID:     object.MCID,
		Email:    object.Email,
	})
	if result.Error != nil {
		return errHandler.DbError(result.Error)
	}

	middleware.RespSuccess(ctx)
	return nil
}
