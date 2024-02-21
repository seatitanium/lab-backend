package backend

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"seatimc/backend/utils"
	"time"
)

func main() {
	dbConf := Conf().Database
	Db, err := sql.Open("mysql", dbConf.User+":"+dbConf.Password+"@"+dbConf.Host+"/"+dbConf.DbName+"?parseTime=true")
	utils.MustPanic(err)
	Db.SetConnMaxLifetime(time.Minute * 3)
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)

	router := gin.New()
	router.Use(middlewares())

}
