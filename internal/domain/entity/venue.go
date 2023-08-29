package entity

import (
	"time"

	"gorm.io/gorm"
)

type Venue struct {
	ID          uint           `json:"id" gorm:"autoIncreament;primaryKey"`
	Name        string         `json:"name"`
	Address     string         `json:"username"`
	Description string         `json:"description"`
	Cost        string         `json:"cost"`
	Status      bool           `json:"status"`
	Days        []Day          `json:"days" gorm:"foreignKey:VenueID"`
	CreateAt    time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt    time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt    gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
