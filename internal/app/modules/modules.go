package modules

import "project-layout/internal/app/modules/user"

type ModuleContainer struct {
	UserModule *user.UserModule
}

type RepositoryContainer struct {
	UserRepository user.UserRepositoryInterface
}
