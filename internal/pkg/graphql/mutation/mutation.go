package mutation

import (
	"project-layout/internal/core/container"
	userContainer "project-layout/internal/services/portal/modules/user/container"

	"github.com/graphql-go/graphql"
)

type MutationFunction func(map[string]*graphql.Object, *container.ModuleContainer) graphql.Fields

func RootMutationType(
	outputTypes map[string]*graphql.Object,
	modules *container.ModuleContainer,
) graphql.Fields {
	mutationFuncs := []MutationFunction{
		userContainer.UserMutations,
	}
}
