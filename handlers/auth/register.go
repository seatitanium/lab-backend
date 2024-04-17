package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/utils"
)

func HandleRegister() gin.HandlerFunc {
	return func(c *gin.Context) {
		conn := utils.GetDBConn()

		var object RegisterRequest
		if err := c.ShouldBindJSON(&object); err != nil {
			utils.RespondNG(c, "Invalid Request Body", "")
			return
		}

		hash, err := utils.GenerateHash(object.Password)
		if err != nil {
			utils.RespondNG(c, "Failed to generate hash for user: "+err.Error(), "")
			return
		}

		result := conn.Create(&utils.Users{
			Username: object.Username,
			Hash:     hash,
			MCID:     object.MCID,
			Email:    object.Email,
		})
		if result.Error != nil {
			utils.RespondNG(c, "Failed to execute some statements: "+result.Error.Error(), "")
			return
		}

		utils.RespondOK(c, "Registered successfully.", "注册成功")
	}
}
