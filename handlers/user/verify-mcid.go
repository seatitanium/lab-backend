package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleMCIDVerify(ctx *gin.Context) *errors.CustomErr {
	playername := ctx.DefaultQuery("playername", "")

	if playername == "" {
		return errors.WrongParam()
	}

	customErr := utils.UpdateUserByMCID(playername, map[string]any{
		"mc_id_verified": true,
	})

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, nil)

	return nil
}
