package utils

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	Conn *gorm.DB
}

var DB *Database
var StatsDB *Database

func Load(dsn string) (error, *gorm.DB) {
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err, nil
	}

	sqlDB, err := conn.DB()
	if err != nil {
		return err, nil
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetConnMaxLifetime(time.Minute * 3)

	return nil, conn
}

func GetDBConn() *gorm.DB {
	return DB.Conn
}

func GetStatsDBConn() *gorm.DB {
	return StatsDB.Conn
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
