package db

import (
	"fmt"

	"driver/conf"

	"github.com/jinzhu/gorm"
)

func getDatabaseConnection() string {
	return conf.MYSQL_USER + ":" + conf.MYSQL_PASSWORD + "@(" + conf.MYSQL_HOST + ":" + conf.MYSQL_PORT + ")/" + conf.MYSQL_DATABASE + "?charset=utf8&parseTime=true"
}

func GetDbInstance() *gorm.DB {
	dbInstance, err := gorm.Open("mysql", getDatabaseConnection())
	if err != nil {
		fmt.Printf("connect database with error: %v\n", err)
		return nil
	}
	return dbInstance
}

func getEcDatabaseConnection() string {
	return conf.MYSQL_USER + ":" + conf.MYSQL_PASSWORD + "@(" + conf.MYSQL_HOST + ":" + conf.MYSQL_PORT + ")/" + conf.MYSQL_EC_DATABASE + "?charset=utf8&parseTime=true"
}

func GetEcDbInstance() *gorm.DB {
	dbInstance, err := gorm.Open("mysql", getEcDatabaseConnection())
	if err != nil {
		fmt.Printf("connect database with error: %v\n", err)
		return nil
	}
	return dbInstance
}
