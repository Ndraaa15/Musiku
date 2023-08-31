package entity

import (
	"time"
)

type Studio struct {
	ID           uint      `json:"id" gorm:"autoIncreament;primaryKey"`
	Name         string    `json:"name"`
	Address      string    `json:"address"`
	Description  string    `json:"description"`
	PricePerHour float64   `json:"price_per_hour"`
	OpenHour     string    `json:"open_hour"`
	Status       bool      `json:"status"`
	Days         []Day     `json:"days" gorm:"foreignKey:StudioID"`
	CreateAt     time.Time `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt     time.Time `json:"update_at" gorm:"autoUpdateTime"`
}
