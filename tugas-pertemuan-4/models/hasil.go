package models

import "time"

type Hasil struct {
	ID        uint `gorm:"primaryKey"`
	TugasID   uint
	Tugas     Tugas `gorm:"foreignKey:TugasID;references:ID"`
	Nilai     uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

// https://gorm.io/docs/conventions.html#Pluralized-Table-Name
func (Hasil) TableName() string {
	return "hasil"
}
