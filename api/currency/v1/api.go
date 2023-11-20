package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func WriteReponse(c *fiber.Ctx, body any, status int) error {
	if body == nil {
		return c.SendStatus(status)
	}
	bytes, err := json.Marshal(body)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	c.Context().Response.SetBodyRaw(bytes)
	c.Context().Response.Header.SetContentType("application/json")
	return c.SendStatus(status)
}
