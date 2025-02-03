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

	res, err := store.SearchCustomer(ctx, "wil")
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range res {
		fmt.Println(v.Name)
	}

	// for range 10000 {
	// 	err = store.SaveClient(ctx, models.Client{
	// 		Name:    gofakeit.Name(),
	// 		Address: gofakeit.Address().Address,
	// 		Phone:   gofakeit.Phone(),
	// 	})
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
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
	// err = store.UpsertProduct(ctx, "Стакан", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// err = store.UpsertProduct(ctx, "Снеговик", 123)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}
