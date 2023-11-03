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

type Route struct {
	URI     string
	Method  string
	Handler fiber.Handler
}

func RegisterRoutes(app *fiber.App, rs []Route) {
	for _, r := range rs {
		switch r.Method {
		case http.MethodGet:
			app.Get(r.URI, r.Handler)
		case http.MethodPost:
			app.Post(r.URI, r.Handler)
		case http.MethodPut:
			app.Put(r.URI, r.Handler)
		case http.MethodPatch:
			app.Patch(r.URI, r.Handler)
		case http.MethodDelete:
			app.Patch(r.URI, r.Handler)
		}
	}
}
