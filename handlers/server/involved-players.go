package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"strconv"
)

func HandleInvolvedPlayers(ctx *gin.Context) *errors.CustomErr {
	termNumber := ctx.DefaultQuery("term", "")

	if termNumber == "" {
		return errors.WrongParam()
	}

	term, err := strconv.Atoi(termNumber)

	if err != nil {
		return errors.WrongParam()
	}

	if term < 7 {
		return errors.TargetNotExist()
	}

	if term > utils.GlobalConfig.ActiveTerm {
		return errors.TargetNotExist()
	}

	if term < 13 {
		history, hErr := utils.GetHistoryTermPlayers()

		if hErr != nil {
			return errors.ServerError(hErr)
		}

		if !utils.HasKey(history, "st"+termNumber) {
			return errors.TargetNotExist()
		}

		players := history["st"+termNumber]

		handlers.RespSuccess(ctx, players)

		return nil
	}

	players, customErr := utils.GetTermPlayers(termNumber)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, players)

	return nil
}
