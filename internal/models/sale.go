package models

import (
	"time"
)

type Sale struct {
	ClientHash string    `bson:"client_hash"`
	Products   []Product `bson:"products"`
	Total      int       `bson:"total"`
	Time       time.Time `bson:"time"`
}
