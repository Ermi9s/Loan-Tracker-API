package routes

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	controller "github.com/Loan-Tracker-API/Loan-Tracker-API/delivery/controller/log"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/middleware"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/repository"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/usecase"
	"github.com/gin-gonic/gin"
)

func NewLogRoute(group *gin.RouterGroup , system_collection database.CollectionInterface , user database.CollectionInterface) {
	
	LogRepository := repository.NewSystemLogsRepository(system_collection)
	LogUsecase := usecase.NewLogUsecase(LogRepository)
	logController := controller.NewLogController(LogUsecase)

	mustBeAdminMiddleware := middleware.RoleBasedAuth(true , user)

	group.GET("/admin/logs",mustBeAdminMiddleware ,logController.GetAllEvents())

}