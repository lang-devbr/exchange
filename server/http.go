package server

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	v1 "github.com/lang-dev/exchange/api/product/v1"
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

	getHandler := v1.NewGetHandler(db)
	app.Get(getHandler.Route, getHandler.Get)

	postHandler := v1.NewPostHandler(db)
	app.Post(postHandler.Route, postHandler.Post)

	log.Printf("app is running at port %d on %s enviroment...", config.Port, config.Env)
	app.Listen(fmt.Sprintf(":%d", config.Port))
}
