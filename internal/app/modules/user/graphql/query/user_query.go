package query

import (
	"github.com/graphql-go/graphql"
)

func UserQueryType() graphql.Fields {
	return graphql.Fields{
		"users": &graphql.Field{},
	}
}
