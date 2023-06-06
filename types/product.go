package types

import "fmt"

const minProductNameLen = 3

type Product struct {
	SKU  string `bson:"sku" json:"sku"`
	Name string `bson:"name" json:"name"`
	Slug string `bson:"slug" json:"slug"`
}

type CreatProductRequest struct {
	SKU  string `json:"sku"`
	Name string `json:"name"`
}

func NewProductFromRequest(req *CreatProductRequest) (*Product, error) {
	if err := validateCreateProductRequest(req); err != nil {
		return nil, err
	}
	return &Product{
		SKU:  req.SKU,
		Name: req.Name,
	}, nil
}

func validateCreateProductRequest(req *CreatProductRequest) error {
	if len(req.Name) < minProductNameLen {
		return fmt.Errorf("")
	}
	return nil
}
