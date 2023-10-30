package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/lang-dev/exchange/api/product/v1"
	"github.com/lang-dev/exchange/config"
)

func StartHTTP() {
	config.Load()

	app := fiber.New(fiber.Config{
		DisableStartupMessage:    false,
		EnableSplittingOnParsers: true,
		ReadTimeout:              config.ReadTimeoutMillis,
		WriteTimeout:             config.WriteTimeoutMillis,
	})

	getHandler := v1.NewGetHandler()
	app.Get(getHandler.Route, getHandler.Get)

	fmt.Printf("app is running at port %d on %s enviroment...", config.Port, config.Env)
	app.Listen(fmt.Sprintf(":%d", config.Port))
}
