package main

import (
	"context"
	"fmt"
	"log"

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
	// err = store.UpsertProduct(ctx, "Стакан", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = store.UpsertProduct(ctx, "Снеговик", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = store.UpsertProduct(ctx, "Стакан", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = store.UpsertProduct(ctx, "Снеговик", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = store.UpsertProduct(ctx, "Стакан", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = store.UpsertProduct(ctx, "Снеговик", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	fmt.Println(store.StorageContent(ctx))

	err = store.RemoveProduct(ctx, "Снеговик", 1087)

	fmt.Println(store.StorageContent(ctx))
}
