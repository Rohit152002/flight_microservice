package controller

import (
	"net/http"
	"user-services/models"
	"user-services/services"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type UserController struct {
	UserServices *services.UserServices
	Logger       *zap.Logger
}

// CreateUserController godoc
// @Summary Create a new user
// @Description Create a new user with the input payload
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User"
// @Success 201 {object} models.User
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /users [post]
func (c *UserController) CreateUserController(ctrl *gin.Context) {
	var user models.User
	if err := ctrl.ShouldBindJSON(&user); err != nil {
		ctrl.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid input"})
		return
	}

	createdUser, err := c.UserServices.CreateUserServices(user)
	if err != nil {
		ctrl.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid input"})
		return
	}
	ctrl.JSON(http.StatusCreated, gin.H{"data": createdUser})
}

// GetAllUserController godoc
// @Summary Get all Users
// @Description Get all users
// @Tags users
// @Accept json
// @Produce json
// @Success 200 {array} models.User
// @Failure 400 {object} map[string]string "Invalid input"
// @Router /users [get]
func (c *UserController) GetAllUserController(ctrl *gin.Context) {
	users, err := c.UserServices.GetAllUserServices()
	if err != nil {
		ctrl.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Invalid input"})
		return
	}
	ctrl.JSON(http.StatusOK, gin.H{"data": users})
}
