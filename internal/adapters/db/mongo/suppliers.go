package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"warehouse/internal/models"
)

// SaveSupplier saves new Supplier.
func (s *Store) SaveSupplier(ctx context.Context, Customer models.Supplier) error {

	_, err := s.supplCl.InsertOne(ctx, Customer)
	if err != nil {
		return err
	}

	return nil
}

// SearchSupplier returns an array of suppliers found in the database according to a given name.
func (s *Store) SearchSupplier(ctx context.Context, name string) ([]models.Supplier, error) {
	filter := bson.D{
		{Key: "name", Value: bson.D{
			{Key: "$regex", Value: name},
			{Key: "$options", Value: "i"},
		}},
	}

	cursor, err := s.supplCl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var res []models.Supplier

	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
