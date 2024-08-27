package controller

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/domain"
	"github.com/gin-gonic/gin"
)

type LogController struct {
	uc domain.SystemLogsUsecase
}

func NewLogController(uc domain.SystemLogsUsecase) *LogController {
	return &LogController{
		uc: uc,
	}
}

func (controller *LogController) GetAllEvents() gin.HandlerFunc {
	return func(c *gin.Context) {
		logs, err := controller.uc.GetAllEvents()
		if err != nil {
			c.IndentedJSON(400, gin.H{"error": err.Error()})
			return
		}
		c.IndentedJSON(200, gin.H{"logs": logs})
	}
}