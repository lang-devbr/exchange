package v1

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lang-dev/exchange/internal/infra"
	"gorm.io/gorm"
)

type GetHandler struct {
	Route        string
	CurrencyRepo infra.CurrencyRepository
}

func NewGetHandler(db *gorm.DB) *GetHandler {
	return &GetHandler{
		Route:        "/v1/converter/:origin/value",
		CurrencyRepo: *infra.NewCurrencyRepository(db),
	}
}

func (h *GetHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return WriteReponse(c, "id cannot be empty", http.StatusBadRequest)
	}

	p, err := h.CurrencyRepo.FindAll()
	for i := 0; i < len(p); i++ {
		//chamar a api que converte

	}

	if err != nil {
		return WriteReponse(c, err.Error(), http.StatusBadRequest)
	}

	return WriteReponse(c, p, http.StatusOK)
}
