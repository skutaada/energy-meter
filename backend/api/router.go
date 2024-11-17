package api

import (
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/energy", EnergyGetHandler)
	router.GET("/energy/current", EnergyGetCurrentHandler)
	router.GET("/energy/:date", EnergyGetDate)

	return router
}
