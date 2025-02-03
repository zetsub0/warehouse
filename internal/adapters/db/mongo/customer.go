package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"warehouse/internal/models"
)

// SaveCustomer saves new Customer.
func (s *Store) SaveCustomer(ctx context.Context, Customer models.Customer) error {

	_, err := s.customerCl.InsertOne(ctx, Customer)
	if err != nil {
		return err
	}

	return nil
}

// SearchCustomer returns an array of customers found in the database according to a given name.
func (s *Store) SearchCustomer(ctx context.Context, name string) ([]models.Customer, error) {
	filter := bson.D{
		{Key: "name", Value: bson.D{
			{Key: "$regex", Value: name},
			{Key: "$options", Value: "i"},
		}},
	}

	cursor, err := s.customerCl.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	var res []models.Customer

	if err = cursor.All(ctx, &res); err != nil {
		return nil, err
	}
	return res, nil
}
