package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleFirstLogin(ctx *gin.Context) *errors.CustomErr {
	mcid := ctx.DefaultQuery("playername", "")
	tag := ctx.DefaultQuery("tag", "")

	if mcid == "" {
		return errors.WrongParam()
	}

	first, customErr := utils.GetFirstLoginRecord(mcid, tag)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, first)

	return nil
}
