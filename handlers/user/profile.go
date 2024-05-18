package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errHandler"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
)

func HandleUserProfile(ctx *gin.Context) *errHandler.CustomErr {
	username := ctx.Param("username")

	if username == "" {
		return errHandler.TargetNotExist()
	}

	user, customErr := utils.GetUserByUsername(username)

	if customErr != nil {
		return customErr
	}

	middleware.RespSuccess(ctx, PublicUser{
		Id:        user.Id,
		Username:  user.Username,
		Nickname:  user.Nickname,
		Email:     user.Email,
		MCID:      user.MCID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	})

	return nil
}
