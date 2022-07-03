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
	Create() (*domain.User, error)
}

// Handler represents handler type
type Handler struct {
	e *echo.Echo

	appPort string
}

// New instance the handler
func New(e *echo.Echo, port string) (*Handler, error) {
	if e == nil || len(port) == 0 {
		return nil, errors.New("[New]: unable to new handler, please check handler arguments")
	}

	return &Handler{
		e:       e,
		appPort: port,
	}, nil
}

// RunServer run the listen and server the service
func (h *Handler) RunServer() error {
	err := h.initRoutes()
	if err != nil {
		return err
	}

	s := http.Server{
		Addr:    ":" + h.appPort,
		Handler: h.e,
	}

	log.Printf("[Handler.RunServer]: initial the service on port %s ... \n", h.appPort)

	err = s.ListenAndServe()
	if err != nil {
		return err
	}

	return nil
}
