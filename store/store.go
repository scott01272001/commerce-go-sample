package store

import "github.com/scott/ggcommerce/types"

// client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
// 	if err != nil {
// 		panic(err)
// 	}

type ProductStorer interface {
	Insert(*types.Product) error
	GetByID(string) (*types.Product, error)
}
