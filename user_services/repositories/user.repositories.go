package repositories

import (
	"user-services/models"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db     *gorm.DB
	Logger *zap.Logger
}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	r.Logger.Info("repositories:= task new created")
	if err := r.Db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	if err := r.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
