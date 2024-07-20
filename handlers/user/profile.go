package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleUserProfile(ctx *gin.Context) *errors.CustomErr {
	username := ctx.DefaultQuery("username", "")

	if username == "" {
		return errors.TargetNotExist()
	}

	user, customErr := utils.GetUserByUsername(username)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, PublicUser{
		Id:           user.Id,
		Username:     user.Username,
		Nickname:     user.Nickname,
		Email:        user.Email,
		MCID:         user.MCID,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
		MCIDVerified: user.MCIDVerified,
	})

	return nil
}
