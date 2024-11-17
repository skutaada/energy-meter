package api

import (
	"backend/db"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func EnergyGetHandler(c *gin.Context) {
	res, err := db.GetLatestEnergy()
	if err != nil {
		c.Error(err)
	}
	c.JSON(http.StatusOK, res)
}

func EnergyGetCurrentHandler(c *gin.Context) {
	res := db.GetCurrentEnergySpot()
	c.JSON(http.StatusOK, res)
}

func EnergyGetDate(c *gin.Context) {
	date, err := time.Parse(time.DateOnly, c.Param("date"))
	if err != nil {
		fmt.Printf("Error parsing time %s", err)
	}
	res := db.GetDayEnergySpot(date)
	c.JSON(http.StatusOK, res)
}
