package user

import (
	"project-layout/internal/app/modules"
	"project-layout/internal/app/modules/user/graphql/output"
	"project-layout/internal/app/modules/user/graphql/query"

	"github.com/graphql-go/graphql"
)

type UserSchema struct {
	Output   map[string]*graphql.Object
	Query    graphql.Fields
	Mutation graphql.Fields
}

func NewUserSchema(
	repositories modules.RepositoryContainer,
) UserSchema {
	return UserSchema{
		Output: output.UserOutputType(),
		Query:  query.UserQueryType(),
		// Mutation: mutation.UserMutationType(repositories),
	}
}
