package common

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"seatimc/backend/handlers/auth"
	"seatimc/backend/handlers/ecs"
	"seatimc/backend/middleware"
	"seatimc/backend/utils"
	"time"
)

type Router struct {
	Port   int
	Router *gin.Engine
}

func (r *Router) Init() {
	log.Println("Using Gin " + gin.Version)

	r.Router = gin.Default()

	r.Router.NoMethod(handleNotFound)
	r.Router.NoRoute(handleNotFound)

	r.Router.Use(cors.New(cors.Config{
		AllowOrigins:     utils.GlobalConfig.AllowedOrigins,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Token"},
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
	ecsGroup.DELETE("delete", wrapper(ecs.HandleDeleteInstance))

	err := r.Router.Run(fmt.Sprintf(":%d", r.Port))
	utils.MustPanic(err)
}
