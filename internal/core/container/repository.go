package container

import "project-layout/internal/modules/user"

type RepositoryContainer struct {
	UserRepository user.UserRepositoryInterface
}
