package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/utils"
)

func HandleRegister(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var object RegisterRequest
		if err := c.ShouldBindJSON(&object); err != nil {
			utils.RespondNG(c, "Invalid Request Body", "")
			return
		}

		tx := db.MustBegin()
		hash, err := utils.GenerateHash(object.Password)
		if err != nil {
			utils.RespondNG(c, "Failed to generate hash for user: "+err.Error(), "")
			return
		}

		_, err = tx.Exec(
			"INSERT INTO `seati_users` (`username`, `hash`, `mcid`, `email`) VALUES (?, ?, ?, ?)",
			object.Username, hash, object.MCID, object.Email,
		)

		if err != nil {
			utils.RespondNG(c, "Failed to execute some statements: "+err.Error(), "")
			return
		}

		err = tx.Commit()

		if err != nil {
			utils.RespondNG(c, "Failed to save user: "+err.Error(), "")
			return
		}

		utils.RespondOK(c, "Registered successfully.", "注册成功")
	}
}
