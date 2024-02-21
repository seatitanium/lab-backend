package lab_backend

import (
	"github.com/gin-gonic/gin"
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
	if currentOrigin := ctx.Request.Header.Get("Origin"); !slices.Contains(Conf().AllowedOrigins, currentOrigin) {
		Respond(ctx, false, "Not supported origin", nil)
	}
}

func midfuncAccessControl(ctx *gin.Context) {
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*"+Conf().Domain)
	ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET")
}

func midfuncTokenCheck(ctx *gin.Context) {
	checkErr := CheckJWT(ctx.Request.Header.Get("Token"))

	if checkErr != nil {
		Respond(ctx, false, "Invalid credentials", nil)
	}
}
