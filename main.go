package main

import (
	"context"

	"github.com/scott/ggcommerce/api"
	"github.com/scott/ggcommerce/store"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/anthdm/weavebox"
)

func main() {
	app := weavebox.New()
	adminRoute := app.Box("/admin")

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	productStore := store.NewMongoProductStore(client)
	productHandler := api.NewProductHandler(productStore)

	adminRoute.Get("/product", productHandler.HandlerGetProduct)
	adminRoute.Post("/product", productHandler.HandlerPostProduct)

	app.Serve(3001)
}
