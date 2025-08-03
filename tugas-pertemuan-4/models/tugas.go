package models

import "time"

type Tugas struct {
	ID          uint      `gorm:"primaryKey"`
	Judul       string    `gorm:"type:varchar(255);not null"`
	Deskripsi   string    `gorm:"type:text;not null"`
	MahasiswaId uint      `gorm:"not null"`
	Mahasiswa   Mahasiswa `gorm:"foreignKey:MahasiswaId;references:ID"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// https://gorm.io/docs/conventions.html#Pluralized-Table-Name
func (Tugas) TableName() string {
	return "tugas"
}
