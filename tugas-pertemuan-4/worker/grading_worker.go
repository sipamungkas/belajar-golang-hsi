package worker

import (
	"fmt"
	"math/rand"
	"sync"
	"tugas-pertemuan-4/models"

	"gorm.io/gorm"
)

func GradingWorker(db *gorm.DB, task models.Tugas, taskChan chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	randValue := uint(rand.Intn(100))

	db.Create(&models.Hasil{
		TugasID: task.ID,
		Nilai:   randValue,
	})

	msg := fmt.Sprintf("Nilai %v diberikan ke '%s' untuk tugas '%s'", randValue, task.Mahasiswa.Nama, task.Judul)

	taskChan <- msg
}
