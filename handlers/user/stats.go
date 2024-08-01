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
	tag := ctx.DefaultQuery("tag", utils.GetActiveTerm().Tag)

	if username == "" && playername == "" {
		return errors.TargetNotExist()
	}

	target, customErr := utils.GetPlayernameByDoubleProvision(username, playername)

	if customErr != nil {
		return customErr
	}

	playtimeRecord, customErr := utils.GetPlaytimeRecord(target, tag)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, playtimeRecord)

	return nil
}

func HandleUserLoginRecords(ctx *gin.Context) *errors.CustomErr {
	username := ctx.DefaultQuery("username", "")
	playername := ctx.DefaultQuery("playername", "")
	tag := ctx.DefaultQuery("tag", utils.GetActiveTerm().Tag)

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

	target, customErr := utils.GetPlayernameByDoubleProvision(username, playername)

	if customErr != nil {
		return customErr
	}

	loginRecords, customErr := utils.GetLoginRecordsByName(target, tag, offset, limit)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, loginRecords)

	return nil
}

func HandleUserLoginRecordTotalCount(ctx *gin.Context) *errors.CustomErr {
	username := ctx.DefaultQuery("username", "")
	playername := ctx.DefaultQuery("playername", "")
	tag := ctx.DefaultQuery("tag", utils.GetActiveTerm().Tag)

	if username == "" && playername == "" {
		return errors.TargetNotExist()
	}

	target, customErr := utils.GetPlayernameByDoubleProvision(username, playername)

	if customErr != nil {
		return customErr
	}

	loginRecords, customErr := utils.GetLoginRecordCount(target, tag)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, loginRecords)

	return nil
}
