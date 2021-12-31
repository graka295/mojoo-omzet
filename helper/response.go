package helper

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// Response contact
type Response interface {
	InternalServerError(c echo.Context, message string) error
	Ok(c echo.Context, data interface{}) error
	BadRequest(c echo.Context, data interface{}) error
	Unauthorized(c echo.Context, data interface{}) error
	Notfound(c echo.Context, message string) error
	Forbidden(c echo.Context, message string) error
}

// ResponseImpl struct for config
type ResponseImpl struct {
}

// NewResponse for init
func NewResponse() Response {
	return &ResponseImpl{}
}

type sharedResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func (x ResponseImpl) InternalServerError(c echo.Context, message string) error {
	return c.JSON(http.StatusInternalServerError, sharedResponse{
		Status:  500,
		Message: message,
	})
}

func (x ResponseImpl) Notfound(c echo.Context, message string) error {
	return c.JSON(http.StatusNotFound, sharedResponse{
		Status:  404,
		Message: message,
	})
}

func (x ResponseImpl) Ok(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, sharedResponse{
		Status:  200,
		Message: "Success",
		Data:    data,
	})
}

func (x ResponseImpl) BadRequest(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusBadRequest, sharedResponse{
		Status:  400,
		Message: "Badrequest",
		Data:    data,
	})
}

func (x ResponseImpl) Unauthorized(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusUnauthorized, sharedResponse{
		Status:  401,
		Message: "Unauthorized",
		Data:    data,
	})
}

func (x ResponseImpl) Forbidden(c echo.Context, message string) error {
	return c.JSON(http.StatusForbidden, sharedResponse{
		Status:  403,
		Message: message,
	})
}
