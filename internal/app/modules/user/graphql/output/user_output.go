package output

import "github.com/graphql-go/graphql"

func UserOutputType() map[string]*graphql.Object {
	outputTypes := make(map[string]*graphql.Object)

	for _, graphqlType := range []*graphql.Object{} {
		outputTypes[graphqlType.Name()] = graphqlType
	}

	return outputTypes
}
