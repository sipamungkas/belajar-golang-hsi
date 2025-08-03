package main

import (
	"fmt"
	"log"
	"tugas-pertemuan-4/config"
	"tugas-pertemuan-4/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	cfg := config.LoadConfig()
	if err := godotenv.Load(); err != nil {
		log.Println(".env not found")
	}

	// this config always create table on "postgres" table, will be check later
	// dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s  TimeZone=Asia/Jakarta",
	// 	cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBPort, cfg.DBSSLMode)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	log.Print("Connected to ", cfg.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	// migration
	db.AutoMigrate(models.Mahasiswa{})
}
