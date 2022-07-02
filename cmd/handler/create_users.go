package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/iamgoangle/gogolf-template/cmd/domain"
)

func createUserHandler(c echo.Context) error {
	users := &domain.UserRequest{}
	err := c.Bind(users)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(users)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users)
}
