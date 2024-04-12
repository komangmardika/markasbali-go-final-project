package models

import (
	"errors"
	"gorm.io/gorm"
)

type Car struct {
	Model
	Name    string `gorm:"not null" json:"name" valid:"required"`
	CarType string `gorm:"not null" json:"car_type" valid:"required"`
	Year    uint   `json:"year" valid:"optional"`
	Uuid    string `gorm:"not null" json:"uuid" valid:"required"`
}

func (car *Car) Upsert(db *gorm.DB) error {
	var b Car
	result := db.Model(Car{}).Where("uuid = ?", car.Uuid).Take(&b)
	if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		result := db.Create(car)
		if result.Error != nil {
			return result.Error
		}
	} else {
		result := db.Model(&b).Updates(car)
		if result.Error != nil {
			return result.Error
		}
	}

	return nil
}
