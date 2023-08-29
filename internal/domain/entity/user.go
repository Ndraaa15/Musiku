package entity

import (
	"time"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID      `json:"id" gorm:"type:char(36);primaryKey"`
	Name     string         `json:"name"`
	Username string         `json:"username"`
	Email    string         `json:"email"`
	Phone    string         `json:"phone" gorm:"type:varchar(13)"`
	CreateAt time.Time      `json:"create_at" gorm:"autoCreateTime"`
	UpdateAt time.Time      `json:"update_at" gorm:"autoUpdateTime"`
	DeleteAt gorm.DeletedAt `json:"delete_at" gorm:"index"`
}
