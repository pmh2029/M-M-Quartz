package container

import "project-layout/internal/pkg/repository"

type RepositoryContainer struct {
	UserRepository repository.UserRepositoryInterface
}
