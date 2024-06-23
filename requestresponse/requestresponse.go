package requestresponse

import (
	"assignment/model"
)

type Order struct {
	ID           string            `json:"id" bson:"_id,omitempty"`
	UserID       string            `json:"user_id" bson:"user_id"`
	RestaurantID string            `json:"restaurant_id" bson:"restaurant_id"`
	Items        []model.OrderItem `json:"items" bson:"items"`
	TotalPrice   float64           `json:"total_price" bson:"total_price"`
	Status       string            `json:"status" bson:"status"`
}
