package models

import (
	"time"
)

type Delivery struct {
	SupplierHash string    `bson:"supplier_hash_hash"`
	Products     []Product `bson:"products"`
	Total        int       `bson:"total"`
	Time         time.Time `bson:"time"`
}
