package output

import (
	"project-layout/internal/pkg/entity"
	"project-layout/internal/pkg/graphql/schema"

	"github.com/graphql-go/graphql"
)

func NewUserType() *graphql.Object {
	return graphql.NewObject(
		graphql.ObjectConfig{
			Name: "user",
			Fields: graphql.FieldsThunk(func() graphql.Fields {
				return graphql.Fields{
					"id": &graphql.Field{
						Type: schema.BigInt,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(entity.User).ID, nil
						},
					},
					"username": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(entity.User).Username, nil
						},
					},
					"email": &graphql.Field{
						Type: graphql.String,
						Resolve: func(p graphql.ResolveParams) (interface{}, error) {
							return p.Source.(entity.User).Email, nil
						},
					},
				}
			}),
		},
	)
}
