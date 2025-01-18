package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"warehouse/internal/adapters/db/mongo"
	"warehouse/internal/config"
)

func main() {

	ctx := context.TODO()

	cfg := config.ParseConfig()

	store, err := mongo.New(ctx, cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}

	sale := Sale{
		SaleDate: time.Now(),
		Client:   "чел с ником кто???",
		Price:    322800,
	}

	res, err := store.AddSale(ctx, sale)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.String())
}

type Sale struct {
	SaleDate time.Time
	Client   string
	Price    int
}
