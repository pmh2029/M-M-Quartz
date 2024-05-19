package validator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

type JsonSchemaValidator struct {
	basePath string
	schemas  map[string]*gojsonschema.Schema // load all schemas file from the base path and store them in the schemas map
}

// NewJsonSchemaValidator creates a new instance of JsonSchemaValidator
func NewJsonSchemaValidator() (*JsonSchemaValidator, error) {
	pwdPath, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	// Replace backslashes with forward slashes for Windows compatibility
	pwdPath = strings.ReplaceAll(pwdPath, "\\", "/")

	// Create a new instance of JsonSchemaValidator
	validator := &JsonSchemaValidator{
		basePath: pwdPath,
		schemas:  make(map[string]*gojsonschema.Schema),
	}

	// Load all schemas from the base path into the schemas map
	// must mount from host to container if using docker compose
	// otherwise it should be ""
	err = validator.loadDirSchemas(os.Getenv("SCHEMA_PATH"))
	if err != nil {
		return nil, err
	}

	return validator, nil
}

// loadDirSchemas loads all schemas from the specified directory path into the schemas map
func (validator *JsonSchemaValidator) loadDirSchemas(path string) error {
	// Walk through all files and directories in the specified directory path
	err := filepath.Walk(validator.basePath+path, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Check if the current file or directory is a directory
		if f.IsDir() {
			return nil
		}

		// Check if the file name ends with.json
		if !strings.HasSuffix(f.Name(), ".json") {
			return nil
		}

		// Get the schema file path
		schemaPath := "file://" + strings.ReplaceAll(path, "\\", "/")

		// Create a new reference loader for the schema file path
		schemaLoader := gojsonschema.NewReferenceLoader(schemaPath)

		// Create a new schema from the schema loader
		schema, err := gojsonschema.NewSchema(schemaLoader)
		if err != nil {
			return err
		}

		// Add the schema to the schemas map using the file name as the key
		validator.schemas[f.Name()] = schema
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

// Validate validates the specified data against the specified schema file
func (validator *JsonSchemaValidator) Validate(schemaFile string, data interface{}) (*gojsonschema.Result, error) {
	// Get the schema from the schemas map
	schema, schemaExists := validator.schemas[schemaFile]
	if !schemaExists {
		return nil, fmt.Errorf("schema %v not found", schemaFile)
	}

	// Create a new Go loader for the data
	dataLoader := gojsonschema.NewGoLoader(data)

	// Validate the data against the schema
	result, err := schema.Validate(dataLoader)
	if err != nil {
		return nil, err
	}

	// Check if there are any validation errors
	if result.Valid() {
		return nil, nil
	}

	// Return the validation errors
	return result, nil
}

func (validator *JsonSchemaValidator) GetErrorDetails(
	result gojsonschema.ResultError,
) map[string]interface{} {
	return map[string]interface{}{
		"context":     result.Context(),
		"description": result.Description(),
		"details":     result.Details(),
		"field":       result.Field(),
		"type":        result.Type(),
		"value":       result.Value(),
	}
}

func (validator *JsonSchemaValidator) GetErrorField(
	result gojsonschema.ResultError,
) string {
	field := result.Field()
	errorDetails := result.Details()
	if property, propertyExists := errorDetails["property"]; propertyExists {
		if propertyString, propertyIsString := property.(string); propertyIsString {
			field = propertyString
		}
	}

	return field
}

func (validator *JsonSchemaValidator) GetCustomErrorMessage(
	result gojsonschema.ResultError,
) string {
	return "add custom error for format rule"
}
