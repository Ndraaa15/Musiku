package entity

import (
	"time"

	"gorm.io/gorm"
)

type Instrument struct {
	ID            uint           `json:"id" gorm:"autoIncreament;primaryKey"`
	Name          string         `json:"name"`
	Rent          float64        `json:"rent"`
	Address       string         `json:"address"`
	Description   string         `json:"description"`
	Spesification string         `json:"spesification"`
	Status        bool           `json:"status"`
	CreateAt      time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt      time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt      gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
