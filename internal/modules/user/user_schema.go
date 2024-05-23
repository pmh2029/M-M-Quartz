package user

import (
	"project-layout/internal/modules/user/graphql/mutation"
	"project-layout/internal/modules/user/graphql/output"
	"project-layout/internal/modules/user/graphql/query"

	"github.com/graphql-go/graphql"
)

func UserOutputType() map[string]*graphql.Object {
	outputTypes := make(map[string]*graphql.Object)

	for _, graphqlType := range []*graphql.Object{
		output.NewUserType(),
	} {
		outputTypes[graphqlType.Name()] = graphqlType
	}

	return outputTypes
}

func UserQueryType(
	outputTypes map[string]*graphql.Object,
) graphql.Fields {
	return graphql.Fields{
		"user": query.UserQueryType(),
	}
}

func UserMutationType(
	outputTypes map[string]*graphql.Object,
) graphql.Fields {
	return graphql.Fields{
		"abc": mutation.UserMutationType(),
	}
}
