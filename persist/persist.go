package persist

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var db *gorm.DB

func InitDB(dbUrl string) {
	var err error
	db, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Database connection established.")
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized. Call InitDB first.")
	}
	return db
}
