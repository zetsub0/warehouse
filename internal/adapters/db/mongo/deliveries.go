package mongo

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"warehouse/internal/models"
)

// AddDelivery saves delivery to database
func (s *Store) AddDelivery(ctx context.Context, delivery models.Delivery) error {
	_, err := s.delivCl.InsertOne(ctx, delivery)
	if err != nil {
		return err
	}

	return nil
}

// GetSuppliersDeliveries takes on the input hash of suppliers
// and returns his purchases
func (s *Store) GetSuppliersDeliveries(ctx context.Context, hash string) ([]models.Delivery, error) {
	var res []models.Delivery

	filter := bson.M{"_id": hash}

	cur, err := s.delivCl.Find(ctx, filter)
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		var delivery models.Delivery
		if err = cur.Decode(&delivery); err != nil {
			return nil, err
		}

		res = append(res, delivery)
	}

	return res, err
}

// GetDeliveriesByDateInterval takes on the input start and end unix timestamps
// and returns deliveries between these dates.
func (s *Store) GetDeliveriesByDateInterval(ctx context.Context, start time.Time, end time.Time) ([]models.Delivery, error) {
	var res []models.Delivery

	filter := bson.D{
		{"$and",
			bson.A{
				bson.D{{"time", bson.D{{"$gt", start}}}},
				bson.D{{"time", bson.D{{"$lte", end}}}},
			},
		},
	}
	cur, err := s.delivCl.Find(ctx, filter)
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err = cur.Close(ctx)
		if err != nil {
			log.Println(err)
		}
	}(cur, ctx)

	for cur.Next(ctx) {
		var delivery models.Delivery
		if err = cur.Decode(&delivery); err != nil {
			return nil, err
		}

		res = append(res, delivery)
	}

	return res, err
}
