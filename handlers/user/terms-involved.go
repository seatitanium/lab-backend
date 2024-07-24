package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleTermsInvolved(ctx *gin.Context) *errors.CustomErr {
	mcid := ctx.DefaultQuery("playername", "")

	if mcid == "" {
		return errors.WrongParam()
	}

	involved, customErr := utils.GetTermsInvolved(mcid)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, involved)

	return nil
}
