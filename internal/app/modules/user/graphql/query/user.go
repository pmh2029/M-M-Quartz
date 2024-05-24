package query

import (
	"project-layout/internal/pkg/graphql/output"

	"github.com/graphql-go/graphql"
)

func User1QueryType() *graphql.Field {
	return &graphql.Field{
		Type: output.BigInt,
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			return 1, nil
		},
	}
}
