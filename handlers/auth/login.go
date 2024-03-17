package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"seatimc/backend/utils"
	"time"
)

func HandleLogin(db *sqlx.DB) gin.HandlerFunc {
	f := func(ctx *gin.Context) {
		var object LoginRequest
		if err := ctx.ShouldBindJSON(&object); err != nil {
			utils.RespondNG(ctx, "Invalid Request Body", "")
			return
		}

		var user User

		tx := db.MustBegin()

		if err := tx.Get(&user, "SELECT * FROM `seati_users` WHERE `username`=?", object.Username); err != nil {
			utils.RespondNG(ctx, "Unable to bind user: "+err.Error(), "")
			return
		}

		if utils.VerifyHash([]byte(user.Hash), []byte(object.Password)) {
			jwt, err := utils.GenerateJWT(utils.JWTPayload{
				Username:  object.Username,
				UpdatedAt: time.Now().UnixMilli(),
			})

			if err != nil {
				utils.RespondNG(ctx, "Unable to create token: "+err.Error(), "")
				return
			}

			utils.ReturnOK(ctx, "Logged in successfully.", "登录成功", jwt)
		} else {
			utils.RespondNG(ctx, "Incorrect password.", "密码错误")
		}
	}

	return f
}
