package main

import (
	"github.com/gin-gonic/gin"
	"seatimc/backend/utils"
	"slices"
)

func middlewares() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		midfuncCheckOrigin(ctx)
		midfuncAccessControl(ctx)
		midfuncTokenCheck(ctx)
	}
}

func midfuncCheckOrigin(ctx *gin.Context) {
	if !utils.GlobalConfig.EnableConfigWhitelist {
		return
	}

	if currentOrigin := ctx.Request.Header.Get("Origin"); !slices.Contains(utils.GlobalConfig.AllowedOrigins, currentOrigin) {
		utils.Respond(ctx, false, "Not supported origin", "", nil)
	}
}

func midfuncAccessControl(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
}

func midfuncTokenCheck(ctx *gin.Context) {
	if !utils.NeedAuthorize(ctx.HandlerName()) {
		return
	}

	checkErr := utils.CheckJWT(ctx.Request.Header.Get("Token"))

	if checkErr != nil {
		utils.Respond(ctx, false, "Invalid credentials", "", nil)
	}
}
