package web

import (
	"project-layout/internal/pkg/handler"
	"project-layout/pkg/shared/validator"

	"github.com/graphql-go/graphql"
	"github.com/sirupsen/logrus"
)

type WebHTTPHandler struct {
	handler.BaseHTTPHandler
	graphql graphql.Schema
}

func NewHTTPHandler(
	graphql graphql.Schema,
	validator *validator.JsonSchemaValidator,
	logger *logrus.Logger,
) *WebHTTPHandler {
	return &WebHTTPHandler{
		BaseHTTPHandler: handler.BaseHTTPHandler{
			Validator: validator,
			Logger:    logger,
		},
		graphql: graphql,
	}
}
