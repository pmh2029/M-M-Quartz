package container

import (
	"project-layout/internal/core/container"
	"project-layout/internal/services/portal/modules/user/pkg/graphql/mutation"

	"github.com/graphql-go/graphql"
)

func UserMutations(
	outputTypes map[string]*graphql.Object,
	modules *container.ModuleContainer,
) graphql.Fields {
	return graphql.Fields{
		"register": mutation.RegisterMutationType(
			outputTypes,
			modules.UserModule,
		),
	}
}
