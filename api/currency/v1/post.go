package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/lang-dev/exchange/internal/dto"
	"github.com/lang-dev/exchange/internal/entity"
	"github.com/lang-dev/exchange/internal/infra"
	"gorm.io/gorm"
)

type PostHandler struct {
	Route        string
	CurrencyRepo infra.CurrencyInterface
}

func NewPostHandler(db *gorm.DB) *PostHandler {
	return &PostHandler{
		Route:        "/v1/currencies",
		CurrencyRepo: infra.NewCurrencyRepository(db),
	}
}

func (h *PostHandler) Post(c *fiber.Ctx) error {
	body := dto.CurrencyRequest{}
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		return WriteReponse(c, nil, http.StatusBadRequest)
	}
	p, err := entity.NewCurrency(body.Code)
	if err != nil {
		return WriteReponse(c, err.Error(), http.StatusBadRequest)
	}
	err = h.CurrencyRepo.Create(p)
	if err != nil {
		return WriteReponse(c, err, http.StatusInternalServerError)
	}
	return WriteReponse(c, nil, http.StatusCreated)
}
