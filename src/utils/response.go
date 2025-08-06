package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse sends a success response
func SuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// ErrorResponse sends an error response
func ErrorResponse(c *gin.Context, statusCode int, message string, err string) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// ValidationErrorResponse sends a validation error response
func ValidationErrorResponse(c *gin.Context, err error) {
	ErrorResponse(c, http.StatusBadRequest, "Validation failed", err.Error())
}
