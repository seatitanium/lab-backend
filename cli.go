package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"log"
	"seatimc/backend/handlers/auth"
	"seatimc/backend/utils"
	"time"
)

func Boot() {
	log.Println("Starting 🌊Tisea Backend.")

	dbConf := utils.Conf().Database
	log.Println("Initializing database with configuration: (mysql) " + dbConf.User + "@" + dbConf.Host + "/" + dbConf.DbName + "?parseTime=true")
	Db, err := sqlx.Open("mysql", dbConf.User+":"+dbConf.Password+"@"+dbConf.Host+"/"+dbConf.DbName+"?parseTime=true")
	utils.MustPanic(err)
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	log.Println("Using Gin " + gin.Version)
	router := gin.New()
	router.Use(middlewares())

	log.Println("Adding routes")
	authGroup := router.Group("/auth")
	authGroup.POST("register", auth.HandleRegister(Db))
	authGroup.POST("login", auth.HandleLogin(Db))
}
