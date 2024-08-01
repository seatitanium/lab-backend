package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
	"strconv"
)

func handleGetAllInvolvedPlayers(ctx *gin.Context) *errors.CustomErr {
	unique := ctx.DefaultQuery("unique", "false")

	uniqueBool := unique == "true"

	history, hErr := utils.GetHistoryTermPlayers()

	if hErr != nil {
		return errors.ServerError(hErr)
	}

	players := make([]utils.ServerPlayer, 0)

	for _, v := range history {
		players = append(players, v...)
	}

	for i := 13; i <= utils.GlobalConfig.ActiveTerm; i++ {
		dbPlayers, customErr := utils.GetTermPlayers("st" + strconv.Itoa(i))

		if customErr != nil {
			continue
		}

		players = append(players, dbPlayers...)
	}

	if uniqueBool {
		players = utils.Unique(players, func(a utils.ServerPlayer, b utils.ServerPlayer) bool {
			return a.UUID == b.UUID
		})
	}

	handlers.RespSuccess(ctx, players)

	return nil
}

func HandleInvolvedPlayers(ctx *gin.Context) *errors.CustomErr {
	termNumber := ctx.DefaultQuery("term", "")

	if termNumber == "" {
		return handleGetAllInvolvedPlayers(ctx)
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

	players, customErr := utils.GetTermPlayers("st" + termNumber)

	if customErr != nil {
		return customErr
	}

	handlers.RespSuccess(ctx, players)

	return nil
}
