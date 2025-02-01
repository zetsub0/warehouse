package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"warehouse/internal/models"
)

// StorageContent returns  all products from storage with their quantities.
func (s *Store) StorageContent(ctx context.Context) ([]models.ProductCount, error) {
	var res []models.ProductCount

	cur, err := s.storageCl.Find(ctx, bson.M{})

	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
		if err != nil {
		}
	}(cur, ctx)

	if err != nil {
		return nil, err
	}

	var raw struct {
		Name  string `bson:"name"`
		Count uint32 `bson:"count"`
	}

	for cur.Next(ctx) {
		if err := cur.Decode(&raw); err != nil {
			return nil, err
		}

		res = append(res, models.ProductCount{
			Product: models.Product{Name: raw.Name},
			Count:   raw.Count,
		})
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}
	return res, err
}

// AddProduct adds quantity for product to storage.
// If the product does not exist, it will be created with the name specified in the "product" parameter.
func (s *Store) AddProduct(ctx context.Context, product string, count int32) error {
	if _, err := s.storageCl.UpdateOne(
		ctx,
		bson.M{"name": product},
		bson.M{
			"$inc": bson.M{
				"count": count,
			},
		},
		options.Update().SetUpsert(true),
	); err != nil {
		return err
	}

	return nil
}

// RemoveProduct subtracts quantity for product from storage.
// If the quantity of an product reaches zero, the product will be deleted.
func (s *Store) RemoveProduct(ctx context.Context, product string, count int32) error {
	if _, err := s.storageCl.UpdateOne(
		ctx,
		bson.M{"name": product},
		bson.M{
			"$inc": bson.M{
				"count": -count,
			},
		},
	); err != nil {
		return err
	}

	if _, err := s.storageCl.DeleteOne(
		ctx,
		bson.M{"name": product, "count": 0},
	); err != nil {
		return err
	}

	return nil
}
