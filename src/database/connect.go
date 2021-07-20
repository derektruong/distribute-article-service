package database

import (
	// "github.com/derektruong/distribute-article-service/src/config"
	// "github.com/derektruong/distribute-article-service/src/model"
	// "fmt"
	// "strconv"

	"gorm.io/driver/mysql"
  	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error
	dbUser := "root"
	dbPass := "123456"
	dbName := "NEWS_APP"

	dsn := dbUser + ":" + dbPass +"@tcp(localhost:3306)/" + dbName + "?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}
}