package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"

	"github.com/iamgoangle/gogolf-template/cmd/domain"
)

//go:generate mockgen -source=handler.go -destination=./mock/service.go UserService

// UserService interface
type UserService interface {
	// Create is a service to perform create a new user
	Create(*domain.User) (*domain.User, error)
}

// Handler represents handler type
type Handler struct {
	Server *echo.Echo
	Port   string

	UserService UserService
}

// New instance the handler
func New(e *echo.Echo, port string, uSrv UserService) (*Handler, error) {
	if e == nil || len(port) == 0 {
		return nil, errors.New("[New]: unable to new handler, please check handler arguments")
	}

	return &Handler{
		Server:      e,
		Port:        port,
		UserService: uSrv,
	}, nil
}

// RunServer run the listen and server the service
func (h *Handler) RunServer() error {
	err := h.initRoutes()
	if err != nil {
		return err
	}

	s := http.Server{
		Addr:    ":" + h.Port,
		Handler: h.Server,
	}

	log.Printf("[Handler.RunServer]: initial the service on port %s ... \n", h.Port)

	err = s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
