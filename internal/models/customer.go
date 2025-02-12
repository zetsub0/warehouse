package models

type Customer struct {
	Name    string `bson:"name"`
	Address string `bson:"address"`
	Phone   string `bson:"phone"`
}
