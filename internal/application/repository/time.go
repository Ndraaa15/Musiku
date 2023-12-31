package repository

import (
	"fmt"

	"github.com/Ndraaa15/musiku/internal/domain/entity"
	"github.com/Ndraaa15/musiku/internal/infrastructure/mysql"
)

func SeedTime(db *mysql.DB) error {
	var dummyData []entity.Time

	for i := 1; i <= 24; i++ {
		hour := fmt.Sprintf("%02d", i)
		time := hour + ":00"
		entity := entity.Time{
			ID:   uint(i),
			Time: time,
		}
		dummyData = append(dummyData, entity)
	}

	if err := db.Create(&dummyData).Error; err != nil {
		return err
	}
	return nil
}
