package routes

import (
	"github.com/Loan-Tracker-API/Loan-Tracker-API/database"
	controller "github.com/Loan-Tracker-API/Loan-Tracker-API/delivery/controller/user"
	enivronment "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/env"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/middleware"
	tokenservice "github.com/Loan-Tracker-API/Loan-Tracker-API/infrastructure/token.service"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/repository"
	"github.com/Loan-Tracker-API/Loan-Tracker-API/usecase"
	"github.com/gin-gonic/gin"
)

func NewUserRoute(group *gin.RouterGroup, users database.CollectionInterface) {
	repo := repository.NewUserRepository(users)
	uc := usecase.NewUserUsecase(repo)
	ctrl := controller.NewUserController(uc)

	key := enivronment.OsGet("SECRETKEY")
	tkokenSev := tokenservice.NewTokenService(key)

	loggedInMiddleware := middleware.LoggedIn(*tkokenSev)
	mustBeAdminMiddleware := middleware.RoleBasedAuth(true , users)
	
	group.GET("/users/profile",loggedInMiddleware, ctrl.GetOneUser())
	group.GET("/admin/users",loggedInMiddleware,mustBeAdminMiddleware, ctrl.GetUsers())
	group.DELETE("admin/users/:id",loggedInMiddleware,mustBeAdminMiddleware, ctrl.DeleteUser())
	group.PUT("/users/:id",loggedInMiddleware, ctrl.UpdatePassword())
}