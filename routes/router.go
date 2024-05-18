package routes

import (
	"github.com/gin-gonic/gin"
	"time-keeping.com/controllers"
)

func Router(router *gin.Engine) {
	router.GET("/attendants/:employeeId", controllers.GetAttendantByID)
	router.GET("/attendants/:employeeId/record-attendances", controllers.GetRecordAttendancesByAttendantID)
}
