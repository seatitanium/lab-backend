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

	dbConf := utils.Conf().Database
	log.Println("Initializing database with configuration: (mysql) " + dbConf.User + "@" + dbConf.Host + "/" + dbConf.DbName + "?parseTime=true")
	Db := utils.GetDb(dbConf)

	log.Println("Using Gin " + gin.Version)
	router := gin.New()
	router.Use(middlewares())

	log.Println("Adding routes")

	versionHandler := func(context *gin.Context) {
		context.String(200, "tisea @ "+utils.Conf().Version)
	}
	// 根目录返回信息
	router.GET("/", versionHandler)
	router.POST("/", versionHandler)

	authGroup := router.Group("/auth")
	authGroup.POST("register", auth.HandleRegister(Db))
	authGroup.POST("login", auth.HandleLogin(Db))

	ecsGroup := router.Group("/ecs")
	ecsGroup.POST("create", ecs.HandleCreateInstance(Db))
	ecsGroup.POST("describe", ecs.HandleDescribeInstance(Db))
	ecsGroup.POST("stop", ecs.HandleStopInstance(Db))
	ecsGroup.POST("start", ecs.HandleStartInstance(Db))
	ecsGroup.POST("reboot", ecs.HandleRebootInstance(Db))

	runErr := router.Run(":" + strconv.Itoa(utils.Conf().BindPort))

	if runErr != nil {
		log.Fatal(runErr.Error())
	}
}

func RunMonitor(monitorName string) {
	Db := utils.GetDb(utils.Conf().Database)

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
			go monitor.RunStoppedInstanceMonitor(Db, time.Second, time.Hour, b)
			break
		}

	default:
		{
			log.Printf("Monitor of name \" %v \" doesn't exist.\n", monitorName)
		}
	}
}
