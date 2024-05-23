package handler

import (
	"encoding/json"
	"errors"
	"net/http"
	"project-layout/pkg/dto"
	"project-layout/pkg/shared/validator"
	"project-layout/pkg/shared/wraperror"

	"github.com/gin-gonic/gin"
	"github.com/graphql-go/graphql/gqlerrors"
	"github.com/sirupsen/logrus"
	"github.com/xeipuuv/gojsonschema"
	"gorm.io/gorm"
)

type BaseHTTPHandler struct {
	Validator *validator.JsonSchemaValidator
	Logger    *logrus.Logger
}

func (h *BaseHTTPHandler) GetInputAsMap(c *gin.Context) (map[string]interface{}, error) {
	contentType := c.ContentType()
	if contentType != "application/json" {
		return nil, errors.New("Content-Type must be application/json")
	}

	// Getting the body as a map
	input := make(map[string]interface{})
	err := c.ShouldBindJSON(&input)
	if err != nil {
		return nil, err
	}

	return input, nil
}

// SetJSONValidationErrorResponse sets the response for a JSON validation error
func (h *BaseHTTPHandler) SetJSONValidationErrorResponse(c *gin.Context, validationResults *gojsonschema.Result) {
	messages := map[string]string{}
	details := make([]map[string]interface{}, 0)
	for _, validationError := range validationResults.Errors() {
		field := h.Validator.GetErrorField(validationError)
		detail := h.Validator.GetErrorDetails(validationError)
		message := h.Validator.GetCustomErrorMessage(validationError)

		messages[field] = message
		details = append(details, detail)
	}

	data := &dto.BaseErrorResponse{
		Code: http.StatusBadRequest,
		Error: &dto.ErrorResponse{
			Details: details,
		},
		Message: "Bad Request",
	}

	c.JSON(http.StatusOK, data)
}

// SetBadRequestErrorResponse sets the response for a bad request error
func (h *BaseHTTPHandler) SetBadRequestErrorResponse(c *gin.Context, err error) {
	data := &dto.BaseErrorResponse{
		Message: "Bad Request",
		Error: &dto.ErrorResponse{
			Details: err.Error(),
		},
	}

	c.JSON(http.StatusBadRequest, data)
}

// SetInternalErrorResponse sets the response for an internal server error
func (h *BaseHTTPHandler) SetInternalErrorResponse(c *gin.Context, err error) {
	data := &dto.BaseErrorResponse{
		Message: "Internal Server Error",
		Error: &dto.ErrorResponse{
			Details: err.Error(),
		},
	}
	// set the response body to the JSON representation of the base error response
	c.JSON(http.StatusInternalServerError, data)
}

func (h *BaseHTTPHandler) SetBadRequestErrorResponseWithCode(c *gin.Context, err error, code int, msg interface{}) {
	data := &dto.BaseErrorResponse{
		Code:    code,
		Message: msg,
		Error: &dto.ErrorResponse{
			Details: err.Error(),
		},
	}

	c.JSON(http.StatusBadRequest, data)
}

func (h *BaseHTTPHandler) SetInternalErrorResponseWithCode(c *gin.Context, err error, code int, msg interface{}) {
	data := &dto.BaseErrorResponse{
		Code:    code,
		Message: msg,
		Error: &dto.ErrorResponse{
			Details: err.Error(),
		},
	}

	c.JSON(http.StatusInternalServerError, data)
}

func (h *BaseHTTPHandler) SetSuccessResponse(c *gin.Context, data interface{}, graphqlQuery string, code int, msg interface{}) {
	c.JSON(http.StatusOK, dto.BaseSuccessResponse{
		Code:    code,
		Data:    data.(map[string]interface{})[graphqlQuery],
		Message: msg,
	})
}

func (h *BaseHTTPHandler) SetGenericErrorResponse(c *gin.Context, finalError error) {
	// Retrieving the original error inside GraphQL's wrapper if there is one
	// If there is none, we keep the error coming from the graphql's engine
	originalError := finalError
	if _, ok := originalError.(gqlerrors.FormattedError); ok {
		err := originalError.(gqlerrors.FormattedError).OriginalError()
		if err != nil {
			originalError = err
		}

		if _, ok := originalError.(*gqlerrors.Error); ok {
			err := originalError.(*gqlerrors.Error).OriginalError
			if err != nil {
				originalError = err
			}
		}
	}

	apiError := &wraperror.ApiDisplayableError{}
	jsonError := &json.SyntaxError{}
	if errors.As(originalError, &apiError) {
		data := dto.BaseErrorResponse{
			Code:    apiError.ErrorCode(),
			Message: finalError.Error(),
			Error: &dto.ErrorResponse{
				Details: apiError.Message(),
			},
		}
		c.JSON(apiError.HttpStatus(), data)
		return
	} else if errors.Is(originalError, gorm.ErrRecordNotFound) || originalError.Error() == gorm.ErrRecordNotFound.Error() {
		data := dto.BaseErrorResponse{
			Error: &dto.ErrorResponse{
				Details: originalError.Error(),
			},
		}
		c.JSON(http.StatusNotFound, data)
		return
	} else if errors.As(originalError, &jsonError) {
		data := dto.BaseErrorResponse{
			Message: "Invalid json",
			Error: &dto.ErrorResponse{
				Details: map[string]interface{}{
					"offset": jsonError.Offset,
					"error":  jsonError.Error(),
				},
			},
		}

		c.JSON(http.StatusBadRequest, data)
		return
	} else {
		h.SetInternalErrorResponse(c, finalError)
		return
	}
}
