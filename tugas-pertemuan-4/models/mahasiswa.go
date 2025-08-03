package models

import "time"

type Mahasiswa struct {
	ID        uint   `gorm:"primaryKey"`
	Nama      string `gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// https://gorm.io/docs/conventions.html#Pluralized-Table-Name
func (Mahasiswa) TableName() string {
	return "mahasiswa"
}
