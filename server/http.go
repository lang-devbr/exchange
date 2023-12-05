package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	currency "github.com/lang-dev/exchange/api/currency/v1"
	product "github.com/lang-dev/exchange/api/product/v1"
	"github.com/lang-dev/exchange/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func StartHTTP() {
	db, err := gorm.Open(sqlite.Open(config.DatabaseName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	app := fiber.New(fiber.Config{
		DisableStartupMessage:    false,
		EnableSplittingOnParsers: true,
		ReadTimeout:              config.ReadTimeoutMillis,
		WriteTimeout:             config.WriteTimeoutMillis,
	})
	app.Use(logger.New())

	getHandlerProduct := product.NewGetHandler(db)
	app.Get(getHandlerProduct.Route, getHandlerProduct.Get)

	postHandlerProduct := product.NewPostHandler(db)
	app.Post(postHandlerProduct.Route, postHandlerProduct.Post)

	getHandlerCurrency := currency.NewGetHandler(db)
	app.Get(getHandlerCurrency.Route, getHandlerCurrency.GetCurrencies)

	postHandlerCurrency := currency.NewPostHandler(db)
	app.Post(postHandlerCurrency.Route, postHandlerCurrency.Post)

	log.Printf("app is running at port %d on %s enviroment...", config.Port, config.Env)
	app.Listen(fmt.Sprintf(":%d", config.Port))
}
