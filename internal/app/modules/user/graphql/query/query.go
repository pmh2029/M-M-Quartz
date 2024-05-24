package query

import (
	"project-layout/internal/app/container"

	"github.com/graphql-go/graphql"
)

func UserQueryType(
	repositories container.RepositoryContainer,
) graphql.Fields {
	return graphql.Fields{
		"users": &graphql.Field{},
	}
}
