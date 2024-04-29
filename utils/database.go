package utils

import (
	"errors"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

var DB *Database

func (db *Database) Load() error {
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
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	DB = &Database{}
	DB.Conn = conn
	return nil
}

func GetDBConn() *gorm.DB {
	return DB.Conn
}

func InitDB() error {
	err := GetDBConn().AutoMigrate(&Users{})
	if err != nil {
		return err
	}

	err = GetDBConn().AutoMigrate(&Ecs{})
	if err != nil {
		return err
	}

	err = GetDBConn().AutoMigrate(&EcsActions{})
	if err != nil {
		return err
	}

	err = GetDBConn().AutoMigrate(&EcsInvokes{})
	if err != nil {
		return err
	}

	return nil
}
