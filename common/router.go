package common

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"seatimc/backend/errors"
	"seatimc/backend/handlers/auth"
	"seatimc/backend/handlers/bss"
	"seatimc/backend/handlers/ecs"
	"seatimc/backend/handlers/server"
	"seatimc/backend/handlers/user"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
	"time"
)

type Router struct {
	Port   int
	Router *gin.Engine
}

type HandlerFunc func(c *gin.Context) *errors.CustomErr

func requestInfo(c *gin.Context) string {
	return fmt.Sprintf("%s %s", c.Request.Method, c.Request.URL.String())
}

func wrapper(handler HandlerFunc) func(c *gin.Context) {
	return func(c *gin.Context) {
		customErr := handler(c)
		if customErr != nil {
			var exception *errors.ApiException
			exception = customErr.Handle()
			exception.Request = requestInfo(c)
			c.JSON(exception.HttpCode, exception)
		}
	}
}

func handleNotFound(ctx *gin.Context) {
	handleErr := errors.NotFound()
	handleErr.Request = requestInfo(ctx)
	ctx.JSON(handleErr.HttpCode, handleErr)
}

func handleVersion(ctx *gin.Context) {
	ctx.String(200, "tisea @ "+utils.GlobalConfig.Version)
}

func handleUnauth(ctx *gin.Context) {
	handleErr := errors.UnAuth()
	handleErr.Request = requestInfo(ctx)
	ctx.JSON(handleErr.HttpCode, handleErr)
}

func (r *Router) Init() {
	log.Println("Using Gin " + gin.Version)

	r.Router = gin.Default()

	r.Router.NoMethod(handleNotFound)
	r.Router.NoRoute(handleNotFound)

	r.Router.Use(cors.New(cors.Config{
		AllowOrigins:     utils.GlobalConfig.AllowedOrigins,
		AllowMethods:     []string{"POST", "GET"},
		AllowCredentials: true,
		MaxAge:           time.Duration(utils.GlobalConfig.Token.Expiration) * time.Minute,
	}))
	r.Router.Use(middleware.TokenCheck)
}

func (r *Router) Run() {
	log.Println("Adding routes")

	r.Router.GET("/", handleVersion)
	r.Router.POST("/", handleVersion)

	authGroup := r.Router.Group("/auth")
	authGroup.POST("register", wrapper(auth.HandleRegister))
	authGroup.POST("login", wrapper(auth.HandleLogin))
	authGroup.GET("check", wrapper(auth.HandleCheck))

	ecsGroup := r.Router.Group("/ecs")
	ecsGroup.GET("create", wrapper(ecs.HandleCreateInstance))
	ecsGroup.GET("describe", wrapper(ecs.HandleDescribeInstance))
	ecsGroup.GET("stop", wrapper(ecs.HandleStopInstance))
	ecsGroup.GET("start", wrapper(ecs.HandleStartInstance))
	ecsGroup.GET("reboot", wrapper(ecs.HandleRebootInstance))
	ecsGroup.GET("deploy-status", wrapper(ecs.HandleGetDeployStatus))
	ecsGroup.DELETE("delete", wrapper(ecs.HandleDeleteInstance))
	ecsGroup.GET("last-invoke", wrapper(ecs.HandleGetInvocationResult))

	bssGroup := r.Router.Group("/bss")
	bssGroup.GET("balance", wrapper(bss.HandleQueryAccountBalance))
	bssGroup.GET("transactions", wrapper(bss.HandleQueryAccountTransactions))

	userGroup := r.Router.Group("/user")
	userGroup.GET("/profile/:username", wrapper(user.HandleUserProfile))
	userGroup.GET("/stats/playtime", wrapper(user.HandleUserPlaytime))
	userGroup.GET("/stats/login", wrapper(user.HandleUserLoginRecord))
	userGroup.GET("/stats/login/count", wrapper(user.HandleUserLoginRecordCount))

	serverGroup := r.Router.Group("/server")
	serverGroup.GET("/online-history", wrapper(server.HandleOnlineHistory))
	serverGroup.GET("/peak-online-history", wrapper(server.HandlePeakOnlineHistory))
	serverGroup.GET("/status", wrapper(server.HandleServerStatus))

	err := r.Router.Run(fmt.Sprintf(":%d", r.Port))
	utils.MustPanic(err)
}
