package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend"
	"seatimc/backend/utils"
)

func HandleRegister(db *sqlx.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var object RegisterRequest
		if err := c.BindJSON(&object); err != nil {
			backend.RespondNg(c, "Invalid Request Body", "", nil)
			return
		}
		tx := db.MustBegin()
		hash, err := utils.GenerateHash(object.Password)
		if err != nil {
			backend.RespondNg(c, "Failed to generate hash for user: "+err.Error(), "", nil)
			return
		}
		_, err = tx.Exec(
			"INSERT INTO `seati_users` (`username`, `hash`, `mcid`, `email`) VALUES ($1, $2, $3, $4)",
			object.Username, hash, object.MCID, object.Email,
		)
		if err != nil {
			backend.RespondNg(c, "Failed to save user: "+err.Error(), "", nil)
			return
		}
		backend.RespondOk(c, "Registered successfully.", "注册成功", nil)
	}
}
