package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
	"os/signal"
	"seatimc/backend/handlers/auth"
	"seatimc/backend/handlers/ecs"
	"seatimc/backend/monitor"
	"seatimc/backend/utils"
	"strconv"
	"syscall"
	"time"
)

func Run() {
	log.Println("Starting 🌊Tisea Backend.")

	log.Println("Using Gin " + gin.Version)
	router := gin.New()
	router.Use(middlewares())

	log.Println("Adding routes")

	versionHandler := func(context *gin.Context) {
		context.String(200, "tisea @ "+utils.GlobalConfig.Version)
	}
	// 根目录返回信息
	router.GET("/", versionHandler)
	router.POST("/", versionHandler)

	authGroup := router.Group("/auth")
	authGroup.POST("register", auth.HandleRegister())
	authGroup.POST("login", auth.HandleLogin())

	ecsGroup := router.Group("/ecs")
	ecsGroup.POST("create", ecs.HandleCreateInstance())
	ecsGroup.POST("describe", ecs.HandleDescribeInstance())
	ecsGroup.POST("stop", ecs.HandleStopInstance())
	ecsGroup.POST("start", ecs.HandleStartInstance())
	ecsGroup.POST("reboot", ecs.HandleRebootInstance())

	runErr := router.Run(":" + strconv.Itoa(utils.GlobalConfig.BindPort))

	if runErr != nil {
		log.Fatal(runErr.Error())
	}
}

func RunMonitor(monitorName string) {

	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	b := make(chan bool)

	// 当接收到中止信号时，将 b 设置为 true
	go func() {
		<-c
		b <- true
	}()

	switch monitorName {
	case "stopped-inst":
		{
			go monitor.RunStoppedInstanceMonitor(time.Second, time.Hour, b)
			break
		}

	default:
		{
			log.Printf("Monitor of name \" %v \" doesn't exist.\n", monitorName)
		}
	}
}
