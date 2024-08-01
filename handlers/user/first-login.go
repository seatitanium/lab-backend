package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleFirstLogin(ctx *gin.Context) *errors.CustomErr {
	playername := ctx.DefaultQuery("playername", "")
	username := ctx.DefaultQuery("username", "")
	tag := ctx.DefaultQuery("tag", "")

	if playername == "" && username == "" {
		return errors.WrongParam()
	}

	target, customErr := utils.GetPlayernameByDoubleProvision(username, playername)

	if customErr != nil {
		return customErr
	}

	// Search in history
	if tag == "" {
		historyLogin := utils.GetHistoryLoginRecord(target)

		if historyLogin != nil {
			handlers.RespSuccess(ctx, historyLogin)
			return nil
		}
	}

	first, customErr := utils.GetFirstLoginRecord(target, tag)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, first)

	return nil
}
