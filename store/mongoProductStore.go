package store

import (
	"github.com/scott/ggcommerce/types"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoProductStore struct {
	client *mongo.Client
	db     *mongo.Database
	coll   string
}

func NewMongoProductStore(c *mongo.Database) *MongoProductStore {
	coll := c.Database("ggcommerce").Collection("products")
	return &MongoProductStore{
		client: c,
		db:     db,
	}
}

func (s *MongoProductStore) Insert(p *types.Product) error {
	_, err := s.db.Collection("products").InsertOne(p)

	return nil
}

func (s *MongoProductStore) GetByID(id string) (*types.Product, error) {
	return nil, nil
}
