package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleMCIDUsage(ctx *gin.Context) *errors.CustomErr {
	mcid := ctx.DefaultQuery("playername", "")

	if mcid == "" {
		return errors.WrongParam()
	}

	usage, customErr := utils.GetMCIDUsage(mcid)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, usage)

	return nil
}
