package infra

import (
	"backend/constant"
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormHandler struct {
	DB *gorm.DB
}

func NewGormHandler() *GormHandler {
	conn, err := connectDB()
	if err != nil {
		panic(err.Error)
	}
	GormHandler := new(GormHandler)
	GormHandler.DB = conn
	return GormHandler
}

func connectDB() (*gorm.DB, error) {

	dbHost := os.Getenv(constant.DBHostEnv)
	dbPort := os.Getenv(constant.DBPortEnv)
	dBName := os.Getenv(constant.DBNameEnv)
	dbUser := os.Getenv(constant.DBUserEnv)
	// dbPassword := os.Getenv(constant.DBPasswordEnv)

	dsn := fmt.Sprintf("%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True&loc=Local&timeout=10s", dbUser, dbHost, dbPort, dBName)
	var db *gorm.DB
	var err error

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	migrateErr := db.AutoMigrate()

	if migrateErr != nil {
		log.Fatal(err)
	}

	return db, nil
}
