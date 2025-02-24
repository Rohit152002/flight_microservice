package routes

import (
	"user-services/controller"
	"user-services/repositories"
	"user-services/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, logger *zap.Logger, db *gorm.DB) {
	user_repo := repositories.UserRepository{Db: db, Logger: logger}
	user_services := services.UserServices{Logger: logger, UserRepository: &user_repo}
	user_controller := controller.UserController{UserServices: &user_services, Logger: logger}

	router.POST("/users", user_controller.CreateUserController)
	router.GET("/users", user_controller.GetAllUserController)
}
