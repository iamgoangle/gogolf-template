package handler

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/iamgoangle/gogolf-template/cmd/domain"
)

func (h *Handler) CreateUserHandler(c echo.Context) error {
	usr := &domain.UserRequest{}
	err := c.Bind(usr)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	err = c.Validate(usr)
	if err != nil {
		return err
	}

	usrService, err := h.UserService.Create(&domain.User{
		Name:  usr.Name,
		Email: usr.Email,
	})
	if err != nil {
		return err
	}

	// PLEASE REMOVE
	log.Println(usrService.Name)
	log.Println(usrService.Email)

	return c.JSON(http.StatusOK, usr)
}
