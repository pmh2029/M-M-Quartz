package output

import (
	"project-layout/internal/services/portal/modules/user/pkg/graphql/output"

	"github.com/graphql-go/graphql"
)

func InitOutput() map[string]*graphql.Object {
	outputTypes := make(map[string]*graphql.Object)
	for _, graphqlType := range []*graphql.Object{
		output.NewUserType(),
	} {
		outputTypes[graphqlType.Name()] = graphqlType
	}

	return outputTypes
}
