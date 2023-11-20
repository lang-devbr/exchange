package v1

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lang-dev/exchange/internal/infra"
	"gorm.io/gorm"
)

type GetHandler struct {
	Route        string
	CurrencyRepo infra.CurrencyInterface
}

func NewGetHandler(db *gorm.DB) *GetHandler {
	return &GetHandler{
		Route:        "/v1/currency/:id",
		CurrencyRepo: infra.NewCurrencyRepository(db),
	}
}

func (h *GetHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return WriteReponse(c, "id cannot be empty", http.StatusBadRequest)
	}

	p, err := h.CurrencyRepo.FindByID(id)
	if err != nil {
		return WriteReponse(c, err.Error(), http.StatusBadRequest)
	}

	return WriteReponse(c, p, http.StatusOK)
}
