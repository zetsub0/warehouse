package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"warehouse/internal/models"
)

// AddSale saves sale to database
func (s *Store) AddSale(ctx context.Context, sale models.Sale) error {
	_, err := s.salesCl.InsertOne(ctx, sale)
	if err != nil {
		return err
	}

	return nil
}

// GetCustomersPurchases takes on the input hash of customer
// and returns his purchases
func (s *Store) GetCustomersPurchases(ctx context.Context, hash string) ([]models.Sale, error) {
	var res []models.Sale

	filter := bson.M{"_id": hash}

	cur, err := s.salesCl.Find(ctx, filter)
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		var sale models.Sale
		if err = cur.Decode(&sale); err != nil {
			return nil, err
		}

		res = append(res, sale)
	}

	return res, err
}

// GetSalesByDateInterval takes on the input start and end unix timestamps
// and returns purchases between these dates.
func (s *Store) GetSalesByDateInterval(ctx context.Context, start time.Time, end time.Time) ([]models.Sale, error) {
	var res []models.Sale

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"time", bson.D{{"$gt", start}}}},
				bson.D{{"time", bson.D{{"$lte", end}}}},
			},
		},
	}
	cur, err := s.salesCl.Find(ctx, filter)
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		var sale models.Sale
		if err = cur.Decode(&sale); err != nil {
			return nil, err
		}

		res = append(res, sale)
	}

	return res, err
}
