package api

import (
	"encoding/json"
	"net/http"

	"github.com/anthdm/weavebox"
	"github.com/scott/ggcommerce/store"
	"github.com/scott/ggcommerce/types"
)

type ProductHandler struct {
	store store.ProductStorer
}

func NewProductHandler(pStore store.ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) HandlerPostProduct(c *weavebox.Context) error {
	productReq := &types.CreatProductRequest{}
	if err := json.NewDecoder(c.Request().Body).Decode(productReq); err != nil {
		return err
	}

	product, err := types.NewProductFromRequest(productReq)
	if err != nil {
		return err
	}

	if err := h.store.Insert(product); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) HandlerGetProduct(c *weavebox.Context) error {
	return c.JSON(http.StatusOK, &types.Product{SKU: "SHOE-1111"})
}
