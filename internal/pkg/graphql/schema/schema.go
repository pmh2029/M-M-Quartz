package schema

import (
	"project-layout/internal/core/container"

	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func NewGraphQLSchema(
	modules container.ModuleContainer,
	db *gorm.DB,
	logger *logrus.Logger,
) (graphql.Schema, error) {
	// outputTypes := output.InitOutput()

	return graphql.NewSchema(graphql.SchemaConfig{})
}
