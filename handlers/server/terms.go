package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleTerms(ctx *gin.Context) *errors.CustomErr {
	terms := utils.GetTerms()

	handlers.RespSuccess(ctx, terms)

	return nil
}
