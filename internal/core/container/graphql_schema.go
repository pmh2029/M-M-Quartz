package container

import (
	"project-layout/internal/modules/user"

	"github.com/graphql-go/graphql"
)

func RootOutputType() map[string]*graphql.Object {
	rootOutputTypes := make(map[string]*graphql.Object)

	for _, outputTypeFunc := range []func() map[string]*graphql.Object{
		user.UserOutputType,
	} {
		for name, graphqlType := range outputTypeFunc() {
			rootOutputTypes[name] = graphqlType
		}
	}

	return rootOutputTypes
}
