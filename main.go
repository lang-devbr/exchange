package main

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/exchange", getHandler)

	app.Listen(":9000")
}

// Exchange api
// - Consultar uma cotação de determinada moeda pela sua sigla na api: https://economia.awesomeapi.com.br/json/last/USD-BRL
// - Salvar essa cotação em um banco de dados
// - Salvar primeiramente no sqlite
// - Alterar para salvar no postgres
// - Conteinerizar o app

func getHandler(c *fiber.Ctx) error {
	r := result{
		Message: "result exchange",
	}

	bytes, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	c.Context().Response.SetBodyRaw(bytes)
	c.Context().Response.Header.SetContentType("application/json")
	c.SendStatus(http.StatusOK)
	return nil
}

type result struct {
	Message string `json:"message"`
}
