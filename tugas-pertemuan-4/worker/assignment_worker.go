package worker

import (
	"fmt"
	"math/rand"
	"tugas-pertemuan-4/models"

	"gorm.io/gorm"
)

var taskList []string = []string{
	"Tugas Pemrograman Goroutine",
	"Tugas Implementasi WaitGroup",
	"Tugas Implementasi Mutex",
	"Tugas Implementasi Channel",
	"Tugas Implementasi WaitGroup"}

// func removeTask(s []string, index int) []string {
// 	return append(s[:index], s[index+1:]...)
// }

func AssignmentWorker(db *gorm.DB, student models.Mahasiswa, msgChan chan<- string) {

	taskCount := len(taskList)

	randomIndex := rand.Intn(taskCount - 1)

	db.Create(&models.Tugas{
		Judul:       taskList[randomIndex],
		Deskripsi:   taskList[randomIndex],
		MahasiswaId: student.ID,
	})

	message := fmt.Sprintf("Tugas '%s' diberikan ke %s", taskList[randomIndex], student.Nama)

	msgChan <- message

}
