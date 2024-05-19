package mutation

import (
	"project-layout/internal/services/portal/modules/user/pkg/repository"

	"github.com/graphql-go/graphql"
)

func RegisterMutationType(
	types map[string]*graphql.Object,
	userRepo repository.UserRepositoryInterface,
) *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "ok", nil
		},
	}
}
