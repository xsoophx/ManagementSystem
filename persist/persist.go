package persist

import (
	"ManagementSystem/persist/models"
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

	migrateUserTable()
	migrateArticleTable()

	log.Println("Database schemes migrated.")
}

func GetDB() *gorm.DB {
	if db == nil {
		log.Fatal("Database not initialized. Call InitDB first.")
	}
	return db
}

func migrateUserTable() {
	err := db.AutoMigrate(&models.UserDbo{})
	if err != nil {
		log.Fatalf("failed to migrate user table: %v", err)
	}
}

func migrateArticleTable() {
	err := db.AutoMigrate(&models.ArticleDbo{})
	if err != nil {
		log.Fatalf("failed to migrate article table: %v", err)
	}
}
