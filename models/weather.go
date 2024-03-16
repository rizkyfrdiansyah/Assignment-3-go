package models

type Weather struct {
	ID    uint `gorm:"primary_key"`
	Water int
	Wind  int
}