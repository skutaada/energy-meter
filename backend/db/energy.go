package db

import (
	"backend/models"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func GetLatestEnergy() (*[]models.EnergySpot, error) {
	res, err := http.Get("https://apis.smartenergy.at/market/v1/price")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	var energyReadings models.EnergyDay
	if err := json.NewDecoder(res.Body).Decode(&energyReadings); err != nil {
		return nil, err
	}
	for _, energy := range energyReadings.Data {
		if err := DB.Create(&energy).Error; err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
			}
		}
	}
	return &energyReadings.Data, nil
}

func GetCurrentEnergySpot() models.EnergySpot {
	var energySpot models.EnergySpot
	now := time.Now()

	interval := now.Minute() / 15 * 15
	targetTime := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		now.Hour(),
		interval,
		0,
		0,
		now.Location(),
	)
	DB.Where("date = ?", targetTime).First(&energySpot)
	return energySpot
}

func GetDayEnergySpot(date time.Time) []models.EnergySpot {
	var spots []models.EnergySpot
	startOfDay := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	endOfDay := startOfDay.Add(24 * time.Hour)
	DB.Where("date BETWEEN ? AND ?", startOfDay, endOfDay).Find(&spots)
	return spots
}
