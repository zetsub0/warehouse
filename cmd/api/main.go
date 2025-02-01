package main

import (
	"context"
	"log"
	"log/slog"

	"warehouse/internal/adapters/db/mongo"
	"warehouse/internal/config"
	"warehouse/internal/logger"
)

func main() {

	ctx := context.TODO()

	cfg := config.ParseConfig()

	myLog := logger.SetupLogger(cfg.Env)
	slog.SetDefault(myLog)

	slog.Error("test err")

	store, err := mongo.New(ctx, cfg.Mongo)
	if err != nil {
		log.Fatal(err)
	}
	_ = store
	// err = store.RemoveProduct(ctx, "asd", 21)

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

	// fmt.Println(store.StorageContent(ctx))

	// err = store.RemoveProduct(ctx, "Снеговик", 1087)

	// fmt.Println(store.StorageContent(ctx))
}
