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
		Route:        "/v1/currency",
		CurrencyRepo: infra.NewCurrencyRepository(db),
	}
}

func (h *GetHandler) GetCurrencies(c *fiber.Ctx) error {
	p, err := h.CurrencyRepo.FindAll()
	if err != nil {
		return WriteReponse(c, err.Error(), http.StatusBadRequest)
	}
	return WriteReponse(c, p, http.StatusOK)
}
