package model

import "time"

type PrecoModel struct {
	ID        string `gorm:"primaryKey"`
	Sku       int `gorm:"not null"`
	Nome      string `gorm:"not null"`
	Valor     float64 `gorm:"not null"`
	CreatedAt *time.Time `gorm:"autoCreateTime"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (PrecoModel) TableName() string {
	return "precos"
}