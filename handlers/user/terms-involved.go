package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleTermsInvolved(ctx *gin.Context) *errors.CustomErr {
	playername := ctx.DefaultQuery("playername", "")
	username := ctx.DefaultQuery("username", "")

	if playername == "" && username == "" {
		return errors.WrongParam()
	}

	target, customErr := utils.GetPlayernameByDoubleProvision(username, playername)

	if customErr != nil {
		return customErr
	}

	involved, customErr := utils.GetTermsInvolved(target)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, involved)

	return nil
}
