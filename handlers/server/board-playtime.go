package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"strconv"
)

func HandleBoardPlaytime(ctx *gin.Context) *errors.CustomErr {
	tag := ctx.DefaultQuery("tag", "")
	limit := ctx.DefaultQuery("limit", "")

	if tag == "" {
		return errors.WrongParam()
	}

	var board []utils.PlaytimeBoard
	var customErr *errors.CustomErr

	if limit == "" {
		board, customErr = utils.GetPlaytimeBoard(tag)
	} else {
		l, err := strconv.Atoi(limit)

		if err != nil {
			return errors.WrongParam()
		}

		if l <= 0 {
			return errors.WrongParam()
		}

		board, customErr = utils.GetPlaytimeBoard(tag, l)
	}

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, board)

	return nil
}
