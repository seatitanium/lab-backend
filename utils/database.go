package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

func (db *Database) open() error {
	var dbConf = GlobalConfig.Database
	if IsStrsHasEmpty(dbConf.Host, dbConf.User, dbConf.Password, dbConf.DBName) || dbConf.Port == 0 {
		return errors.New("connect to database failed due to configuration error")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User, dbConf.Password, dbConf.Host, dbConf.Port, dbConf.DBName)

	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	db.Conn = conn
	return nil
}

func GetDBConn() *gorm.DB {
	db := Database{}
	err := db.open()
	if err != nil {
		log.Fatal(err)
	}
	return db.Conn
}

func InitDB() error {
	conn := GetDBConn()

	err := conn.AutoMigrate(&Users{})
	if err != nil {
		return err
	}

	err = conn.AutoMigrate(&Ecs{})
	if err != nil {
		return err
	}

	err = conn.AutoMigrate(&EcsActions{})
	if err != nil {
		return err
	}

	err = conn.AutoMigrate(&EcsInvokes{})
	if err != nil {
		return err
	}

	return nil
}
