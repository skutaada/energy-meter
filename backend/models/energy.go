package models

import (
	"time"

	"gorm.io/gorm"
)

type EnergyDay struct {
	Unit     string       `json:"unit"`
	Interval uint         `json:"interval"`
	Data     []EnergySpot `json:"data"`
}

type EnergySpot struct {
	gorm.Model
	Date  time.Time `json:"date" gorm:"uniqueIndex"`
	Value float32   `json:"value"`
}
