package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type GetHandler struct {
	Route string
}

func NewGetHandler() *GetHandler {
	return &GetHandler{
		Route: "/v1/product",
	}
}

func (h *GetHandler) Get(c *fiber.Ctx) error {
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
