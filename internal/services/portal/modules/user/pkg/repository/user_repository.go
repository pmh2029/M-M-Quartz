package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface{}

type userRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(
	db *gorm.DB,
	logger *logrus.Logger,
) UserRepositoryInterface {
	return &userRepository{
		db,
		logger,
	}
}
