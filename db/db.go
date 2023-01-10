package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

var (
	HOST    = "localhost"
	PORT    = "5432"
	DB_NAME = "phincon"
	DB_USER = "root"
	DB_PASS = "newpassword"
)

func getLogger() logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
}

func Connect() (err error) {
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", HOST, DB_USER, DB_PASS, DB_NAME, PORT)
	dsn := "root:newpassword@tcp(127.0.0.1:3306)/phincon?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: getLogger(),
	})

	return err
}

func Get() *gorm.DB {
	return db
}
