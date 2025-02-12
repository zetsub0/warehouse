package models

type Supplier struct {
	Name       string                 `bson:"name"`
	Address    string                 `bson:"address"`
	Phone      string                 `bson:"phone"`
	Additional map[string]interface{} `bson:"additional_info"`
}
