package query

import "github.com/graphql-go/graphql"

func UserQueryType() *graphql.Field {
	return &graphql.Field{
		Type: graphql.String,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return "nil", nil
		},
	}
}
