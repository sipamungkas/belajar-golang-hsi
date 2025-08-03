package worker

import (
	"fmt"
	"tugas-pertemuan-4/models"

	"gorm.io/gorm"
)

func AssignmentWorker(db *gorm.DB, student models.Mahasiswa, msgChan chan<- string) {

	taskList := []string{
		"Tugas Pemrograman Goroutine",
		"Tugas Implementasi WaitGroup",
		"Tugas Implementasi Mutex",
		"Tugas Implementasi Channel",
		"Tugas Implementasi WaitGroup"}

	db.Create(&models.Tugas{
		Judul:       taskList[1],
		Deskripsi:   taskList[1],
		MahasiswaId: student.ID,
	})

	message := fmt.Sprintf("Tugas '%s' diberikan ke %s", taskList[1], student.Nama)

	msgChan <- message

}
