package handler

import (
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// CustomValidator represents custom service validator
type CustomValidator struct {
	Validator *validator.Validate
}

// MockValidator type
type MockValidator struct {
	Validator *validator.Validate
}

// Validate validates the struct fields
func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}

// Validate mocking validation
func (m *MockValidator) Validate(i interface{}) error {
	return nil
}
