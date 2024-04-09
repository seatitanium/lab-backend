package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"seatimc/backend/handlers/auth"
	"seatimc/backend/handlers/ecs"
	"seatimc/backend/utils"
	"strconv"
	"time"
)

func Run() {
	log.Println("Starting 🌊Tisea Backend.")

	dbConf := utils.Conf().Database
	log.Println("Initializing database with configuration: (mysql) " + dbConf.User + "@" + dbConf.Host + "/" + dbConf.DbName + "?parseTime=true")
	Db, err := sqlx.Open("mysql", dbConf.User+":"+dbConf.Password+"@tcp("+dbConf.Host+")/"+dbConf.DbName+"?parseTime=true")
	utils.MustPanic(err)
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

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
