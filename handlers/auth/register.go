package auth

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"seatimc/backend"
	"seatimc/backend/utils"
)

func HandleRegister(db *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var object RegisterRequest
		if err := c.BindJSON(&object); err != nil {
			backend.Respond(c, false, "Invalid Request Body", nil)
			return
		}
		stmt, stmtErr := db.Prepare("INSERT INTO `seati_users` (`username`, `hash`, `mcid`, `email`) VALUES (?, ?, ?, ?)")
		defer utils.MustPanic(stmt.Close())
		if stmtErr != nil {
			backend.Respond(c, false, "Unable to create user: "+stmtErr.Error(), nil)
			return
		}
		hash, err := utils.GenerateHash(object.Password)
		if err != nil {
			backend.Respond(c, false, "Failed to generate hash for user: "+err.Error(), nil)
			return
		}
		_, err = stmt.Exec(object.Username, hash, object.MCID, object.Email)
		if err != nil {
			backend.Respond(c, false, "Failed to execute some sql: "+err.Error(), nil)
			return
		}
	}
}
