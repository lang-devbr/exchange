package v1

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
	v1 "github.com/lang-dev/exchange/api"
	"github.com/lang-dev/exchange/internal/dto"
	"github.com/lang-dev/exchange/internal/entity"
	"github.com/lang-dev/exchange/internal/infra"
	"gorm.io/gorm"
)

type ProductHandler struct {
	ProductRepo infra.ProductInterface
}

func NewProductHandler(db *gorm.DB) *ProductHandler {
	return &ProductHandler{
		ProductRepo: infra.NewProductRepository(db),
	}
}

func RegisterProductHandler(app *fiber.App, db *gorm.DB) {
	productHandler := NewProductHandler(db)
	var routes = []v1.Route{
		{
			URI:     "/v1/product/:id",
			Method:  http.MethodGet,
			Handler: productHandler.Get,
		},
		{
			URI:     "/v1/product",
			Method:  http.MethodPost,
			Handler: productHandler.Post,
		},
	}
	v1.RegisterRoutes(app, routes)
}

func (h *ProductHandler) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return v1.WriteReponse(c, "id cannot be empty", http.StatusBadRequest)
	}

	p, err := h.ProductRepo.FindByID(id)
	if err != nil {
		return v1.WriteReponse(c, err.Error(), http.StatusBadRequest)
	}

	return v1.WriteReponse(c, p, http.StatusOK)
}

func (h *ProductHandler) Post(c *fiber.Ctx) error {
	body := dto.ProductRequest{}
	err := json.Unmarshal(c.Body(), &body)
	if err != nil {
		return v1.WriteReponse(c, nil, http.StatusBadRequest)
	}

	p, err := entity.NewProduct(body.Name, body.Price)
	if err != nil {
		return v1.WriteReponse(c, err.Error(), http.StatusBadRequest)
	}

	err = h.ProductRepo.Create(p)
	if err != nil {
		return v1.WriteReponse(c, err, http.StatusInternalServerError)
	}

	return v1.WriteReponse(c, nil, http.StatusCreated)
}
