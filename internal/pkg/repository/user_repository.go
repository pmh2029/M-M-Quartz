package repository

import (
	"context"
	"project-layout/internal/pkg/entity"
	"project-layout/pkg/shared/utils"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepositoryInterface interface {
	Create(
		ctx context.Context,
		attributes map[string]interface{},
	) (entity.User, error)
	TakeByConditions(
		ctx context.Context,
		conditions map[string]interface{},
	) (entity.User, error)
	Update(ctx context.Context,
		user entity.User,
		attributesToUpdate map[string]interface{},
	) (entity.User, error)
}

type UserRepository struct {
	db     *gorm.DB
	logger *logrus.Logger
}

func NewUserRepository(
	db *gorm.DB,
	logger *logrus.Logger,
) UserRepositoryInterface {
	return &UserRepository{
		db:     db,
		logger: logger,
	}
}

func (r *UserRepository) Create(
	ctx context.Context,
	attributes map[string]interface{},
) (entity.User, error) {
	var user entity.User
	err := utils.MapToStruct(attributes, &user)
	if err != nil {
		return entity.User{}, err
	}

	cdb := r.db.WithContext(ctx)
	err = cdb.Create(&user).Error
	return user, err
}

func (r *UserRepository) TakeByConditions(ctx context.Context, conditions map[string]interface{}) (entity.User, error) {
	cdb := r.db.WithContext(ctx)

	var user entity.User
	err := cdb.Where(conditions).Take(&user).Error

	return user, err
}

func (r *UserRepository) Update(ctx context.Context, user entity.User, attributesToUpdate map[string]interface{}) (entity.User, error) {
	err := utils.MapToStruct(attributesToUpdate, &user)
	if err != nil {
		return entity.User{}, err
	}
	cdb := r.db.WithContext(ctx)
	err = cdb.Model(&user).Updates(attributesToUpdate).Error
	return user, err
}
