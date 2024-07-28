package server

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/errors"
	"seatimc/backend/handlers"
	"seatimc/backend/utils"
)

func HandleTermSet(ctx *gin.Context) *errors.CustomErr {
	var termSet struct {
		Tag             string   `json:"tag"`
		Forge           string   `json:"forge"`
		Ram             string   `json:"ram"`
		Java            string   `json:"java"`
		DescriptionList []string `json:"descriptionList,omitempty"`
		Description     string   `json:"description,omitempty"`
	}

	err := utils.GetJsonFromData("current-set.json", &termSet)

	if err != nil {
		return errors.DbError(err)
	}

	if len(termSet.DescriptionList) > 0 {
		var dd string

		for _, d := range termSet.DescriptionList {
			dd += "<p>" + d + "</p>"
		}

		termSet.Description = dd
		termSet.DescriptionList = []string{}
	}

	handlers.RespSuccess(ctx, termSet)

	return nil
}
