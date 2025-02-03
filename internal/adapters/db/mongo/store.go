package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"warehouse/internal/config"
)

type Store struct {
	db        *mongo.Database
	storageCl *mongo.Collection
	pListCl   *mongo.Collection

	supplCl *mongo.Collection
	delivCl *mongo.Collection

	customerCl *mongo.Collection
	salesCl    *mongo.Collection
}

const (
	storageCollection     = "storage"
	productListCollection = "product_list"

	suppliersCollection  = "suppliers"
	deliveriesCollection = "deliveries"

	customersCollection = "customers"
	salesCollection     = "sales"
)

func getDB(ctx context.Context, cfg config.Mongo, maxPoolSize uint64) (*mongo.Database, error) {
	opts := options.Client().SetHosts(cfg.Hosts)

	if cfg.Login != "" && cfg.Password != "" {
		opts.SetAuth(options.Credential{
			AuthSource: cfg.AuthSource,
			Username:   cfg.Login,
			Password:   cfg.Password,
		}).SetMaxPoolSize(maxPoolSize)
	}

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return nil, err
	}

	return client.Database(cfg.DbName), err
}

// New returns Store adapter.
func New(ctx context.Context, cfg config.Mongo) (*Store, error) {
	db, err := getDB(ctx, cfg, cfg.ConnectCount)

	if err != nil {
		return nil, err
	}

	store := &Store{
		db:        db,
		storageCl: db.Collection(storageCollection),
		pListCl:   db.Collection(productListCollection),

		supplCl: db.Collection(suppliersCollection),
		delivCl: db.Collection(deliveriesCollection),

		customerCl: db.Collection(customersCollection),
		salesCl:    db.Collection(salesCollection),
	}

	return store, err
}
