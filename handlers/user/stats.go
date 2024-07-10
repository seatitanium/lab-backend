package user

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"strconv"
)

func HandleUserPlaytime(ctx *gin.Context) *errors.CustomErr {
	username := ctx.DefaultQuery("username", "")
	playername := ctx.DefaultQuery("playername", "")
	tag := ctx.DefaultQuery("tag", "")

	if username == "" && playername == "" {
		return errors.TargetNotExist()
	}

	var targetName string

	if playername == "" {
		user, customErr := utils.GetUserByUsername(username)

		if customErr != nil {
			return customErr
		}

		targetName = user.MCID
	} else {
		targetName = playername
	}

	playtimeRecord, customErr := utils.GetPlaytimeRecord(targetName, tag)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, playtimeRecord)

	return nil
}

func HandleUserLoginRecord(ctx *gin.Context) *errors.CustomErr {
	username := ctx.DefaultQuery("username", "")
	playername := ctx.DefaultQuery("playername", "")
	tag := ctx.DefaultQuery("tag", "")
	offset, err := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	if err != nil {
		return errors.WrongParam()
	}
	limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	if err != nil {
		return errors.WrongParam()
	}

	if username == "" && playername == "" {
		return errors.TargetNotExist()
	}

	var targetName string

	if playername == "" {
		user, customErr := utils.GetUserByUsername(username)

		if customErr != nil {
			return customErr
		}

		targetName = user.MCID
	} else {
		targetName = playername
	}

	loginRecords, customErr := utils.GetLoginRecordsByName(targetName, tag, offset, limit)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, loginRecords)

	return nil
}

func HandleUserLoginRecordCount(ctx *gin.Context) *errors.CustomErr {
	username := ctx.DefaultQuery("username", "")
	playername := ctx.DefaultQuery("playername", "")
	tag := ctx.DefaultQuery("tag", "")

	if username == "" && playername == "" {
		return errors.TargetNotExist()
	}

	var targetName string

	if playername == "" {
		user, customErr := utils.GetUserByUsername(username)

		if customErr != nil {
			return customErr
		}

		targetName = user.MCID
	} else {
		targetName = playername
	}

	loginRecords, customErr := utils.GetLoginRecordCount(targetName, tag)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, loginRecords)

	return nil
}
