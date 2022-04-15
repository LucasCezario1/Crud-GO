package database

import (
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB

func StartDB() {
	str := "host=localhost port=25432 user=admin dbname=movie sslmode=disable password=admin"

	databasse, err = gorm.Open(postgres.Open(str), &gorm.Config{})
	if err != nil {
		log.Fatal("errorrr", err)
	}

	db = database

	config, _ := db.DB()

	config.SetConnMaxIdleTime(time.Hour)
	config.SetMaxIdleConns(10)
	config.SetMaxOpenConns(100)
}

// abrir conexao com o banco
func Getdatabase() *gorm.DB {
	return db
}
