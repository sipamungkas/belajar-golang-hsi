package main

import (
	"fmt"
	"log"
	"sync"
	"tugas-pertemuan-4/config"
	"tugas-pertemuan-4/models"
	"tugas-pertemuan-4/worker"

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
	db.AutoMigrate(models.Tugas{})
	db.AutoMigrate(models.Hasil{})

	var count int64

	// Counting students
	db.Model(&models.Mahasiswa{}).Count(&count)
	log.Println("Current students: ", count)
	// check table count
	if count < 1 {
		newStudents := []*models.Mahasiswa{
			{Nama: "Ragil"},
			{Nama: "Burhan"},
			{Nama: "Udin"},
			{Nama: "Pamung"},
			{Nama: "Ragil P"},
		}
		res := db.Create(newStudents)

		var insertedCount int64
		res.Count(&insertedCount)
		log.Println(insertedCount, "Students Successfully added to table mahasiswa")
	}

	var students []models.Mahasiswa
	res := db.Find(&students)
	if res.Error != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	msgChan := make(chan string, 5)

	for _, student := range students {
		wg.Add(1)
		go worker.AssignmentWorker(db, student, msgChan)
	}

	for range students {
		fmt.Println(<-msgChan)
	}

}
