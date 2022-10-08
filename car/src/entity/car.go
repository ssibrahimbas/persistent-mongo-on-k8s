package entity

import "time"

type Car struct {
	ID           string    `json:"uuid,omitempty" bson:"_id,omitempty"`
	Brand        string    `json:"brand,omitempty" bson:"brand,omitempty"`
	Model        string    `json:"model,omitempty" bson:"model,omitempty"`
	Year         int       `json:"year,omitempty" bson:"year,omitempty"`
	Color        string    `json:"color,omitempty" bson:"color,omitempty"`
	DateOfCreate time.Time `json:"createdAt,omitempty" bson:"date_of_create,omitempty"`
}
