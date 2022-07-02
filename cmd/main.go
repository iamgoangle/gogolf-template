package main

import (
	"log"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/iamgoangle/gogolf-template/cmd/handler"

	"github.com/labstack/echo/v4"
)

var (
	appPort string
	//appName string
	//
	//mongodbHost     string
	//mongodbUsername string
	//mongodbPassword string
	//
	//redisHost     string
	//redisUsername string
	//redisPassword string
)

func initAppConfigs() {
	appPort = os.Getenv("APP_PORT")
	//appName = os.Getenv("APP_NAME")

	log.Print("[initAppConfigs]: initial app config success") // DO NOT USING THIS LOGGER IN PRODUCTION
}

func initRedisConfigs() {

}

func initMongoConfigs() {

}

// ==============================================
// ===== DO NOT USE THIS CODE IN PRODUCTION =====
// ==============================================
func main() {
	initAppConfigs()
	initRedisConfigs()
	initMongoConfigs()

	e := echo.New()
	e.Debug = false
	e.Validator = &handler.CustomValidator{
		Validator: validator.New(),
	}

	h, err := handler.New(e, appPort)
	if err != nil {
		panic(err)
	}

	err = h.RunServer()
	if err != nil {
		panic(err)
	}
}
