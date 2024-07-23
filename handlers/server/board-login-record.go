package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"strconv"
)

func HandleBoardLoginRecord(ctx *gin.Context) *errors.CustomErr {
	tag := ctx.DefaultQuery("tag", "")
	limit := ctx.DefaultQuery("limit", "")

	if tag == "" {
		return errors.WrongParam()
	}

	var board []utils.LoginRecordBoard
	var customErr *errors.CustomErr

	if limit == "" {
		board, customErr = utils.GetLoginRecordBoard(tag)
	} else {
		l, err := strconv.Atoi(limit)

		if err != nil {
			return errors.WrongParam()
		}

		if l <= 0 {
			return errors.WrongParam()
		}

		board, customErr = utils.GetLoginRecordBoard(tag, l)
	}

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, board)

	return nil
}
