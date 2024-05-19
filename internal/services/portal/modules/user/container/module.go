package container

import (
	"project-layout/internal/services/portal/modules/user/pkg/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ModuleContainer struct {
	Repository repository.UserRepositoryInterface
}

func NewUserModuleContainer(
	db *gorm.DB,
	logger *logrus.Logger,
) ModuleContainer {
	return ModuleContainer{
		Repository: repository.NewUserRepository(db, logger),
	}
}
