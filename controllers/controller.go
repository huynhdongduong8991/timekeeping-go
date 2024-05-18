package controllers

import (
	"log"

	"github.com/gin-gonic/gin"
	"time-keeping.com/services"
)

func GetAttendantByID(c *gin.Context) {
	employeeID, exist := c.Params.Get("employeeId")
	if !exist {
		c.JSON(400, gin.H{
			"message": "Empoloyee ID is invalid",
		})
		return
	}

	attendant, err := services.GetAttendantByID(employeeID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, attendant)
}

func GetRecordAttendancesByAttendantID(c *gin.Context) {
	employeeID, exist := c.Params.Get("employeeId")
	if !exist {
		c.JSON(400, gin.H{
			"message": "Empoloyee ID is invalid",
		})
		return
	}

	attendant, err := services.GetRecordAttendancesByAttendantID(employeeID)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
		})
		return
	}

	c.JSON(200, attendant)
}
