package auth

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/utils"
	"time"
)

func HandleLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var object LoginRequest
		if err := ctx.ShouldBindJSON(&object); err != nil {
			utils.RespondNG(ctx, "Invalid Request Body", "")
			return
		}

		user, err := utils.GetUserByUsername(object.Username)

		if err != nil {
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
}
