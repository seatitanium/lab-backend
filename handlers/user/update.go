package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleUpdateUserProfile(ctx *gin.Context) *errors.CustomErr {
	username := ctx.DefaultQuery("username", "")

	if username == "" {
		return errors.TargetNotExist()
	}

	var updateRequest = map[string]any{}

	err := ctx.ShouldBindJSON(&updateRequest)

	if err != nil {
		return errors.WrongParam()
	}

	for k := range updateRequest {
		if utils.NoneMatch(k, "username", "nickname", "mc_id", "email", "password") {
			return errors.WrongParam()
		}
	}

	if utils.HasKey(updateRequest, "mc_id") {

		// Check if already bound
		used, customErr := utils.IsMCIDUsed(updateRequest["mc_id"].(string))

		if customErr != nil {
			return customErr
		}

		if used {
			return errors.DuplicatedMCIDBinding()
		}

		// Reset mc_id_verified state if updating mc_id directly
		updateRequest["mc_id_verified"] = false
	}

	customErr := utils.UpdateUserByUsername(username, updateRequest)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, nil)

	return nil
}
