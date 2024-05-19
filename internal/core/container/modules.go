package container

import (
	userModule "project-layout/internal/services/portal/modules/user/container"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type ModuleContainer struct {
	UserModule userModule.ModuleContainer
}

func NewModuleContainer(
	db *gorm.DB,
	logger *logrus.Logger,
) ModuleContainer {
	return ModuleContainer{
		UserModule: userModule.NewUserModuleContainer(db, logger),
	}
}
