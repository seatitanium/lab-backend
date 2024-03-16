package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/utils"
	"time"
)

func HandleLogin(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var object LoginRequest
		if err := ctx.BindJSON(&object); err != nil {
			utils.RespondNg(ctx, "Invalid Request Body", "", nil)
			return
		}

		var user User

		tx := db.MustBegin()
		err := tx.Select(&user, "SELECT * FROM `seati_users` WHERE username=$1", object.Username)
		if err != nil {
			utils.RespondNg(ctx, "Unable to bind user: "+err.Error(), "", nil)
		}
		if utils.VerifyHash([]byte(object.Password), []byte(user.Hash)) {
			jwt, err := utils.GenerateJWT(utils.JWTPayload{
				Username:  object.Username,
				UpdatedAt: time.Now().UnixMilli(),
			})

			if err != nil {
				utils.RespondNg(ctx, "Unable to create token: "+err.Error(), "", nil)
				return
			}

			utils.RespondOk(ctx, "Logged in successfully.", "登录成功", jwt)
		} else {
			utils.RespondNg(ctx, "Incorrect password.", "", nil)
		}
	}
}
