package model

import "time"

type User struct {
	ID            string `json:"id" bson:"_id,omitempty"`
	Name          string `json:"name" bson:"name" validate:"required"`
	Email         string `json:"email" bson:"email" validate:"required"`
	Password      string `json:"password" bson:"password" validate:"required"`
	Mobile_Number string `json:"mobile_number" bson:"mobile_number" validate:"required"`
}

type Rider struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Name        string   `json:"name" bson:"name" validate:"required"`
	PhoneNumber string   `json:"phone_number" bson:"phone_number" validate:"required"`
	VehicleType string   `json:"vehicle_type" bson:"vehicle_type" validate:"required"`
	Location    Location `json:"location" bson:"location" validate:"required"`
	Orders      []Order  `json:"orders" bson:"orders"`
}

type Location struct {
	Latitude  float64 `json:"latitude" bson:"latitude" validate:"required"`
	Longitude float64 `json:"longitude" bson:"longitude" validate:"required"`
}

type Restaurant struct {
	ID       string     `json:"id" bson:"_id,omitempty"`
	Name     string     `json:"name" bson:"name" validate:"required"`
	Address  string     `json:"address" bson:"address" validate:"required"`
	Cuisine  string     `json:"cuisine" bson:"cuisine" validate:"required"`
	Menu     []MenuItem `json:"menu" bson:"menu" validate:"required"`
	Rating   float64    `json:"rating" bson:"rating"`
	Location Location   `json:"location" bson:"location" validate:"required"`
	Orders   []Order    `json:"orders" bson:"orders"`
}

type MenuItem struct {
	ID          string  `json:"_id" bson:"_id"`
	Name        string  `json:"name" bson:"name"`
	Description string  `json:"description" bson:"description"`
	Price       float64 `json:"price" bson:"price"`
}

type Order struct {
	ID           string      `json:"id" bson:"_id,omitempty"`
	UserID       string      `json:"user_id" bson:"user_id"`
	RiderID      string      `json:"rider_id" bson:"rider_id"`
	Items        []OrderItem `json:"items" bson:"items"`
	TotalPrice   float64     `json:"total_price" bson:"total_price"`
	Status       string      `json:"status" bson:"status"`
	RestaurantID string      `json:"restaurant_id" bson:"restaurant_id"`
	Createdat    time.Time   `json:"_created_at" bson:"_created_at"`
	Updatedat    time.Time   `json:"_updated_at" bson:"_updated_at"`
}

type OrderItem struct {
	MenuItemID string `json:"menu_item_id" bson:"menu_item_id"`
	Quantity   int    `json:"quantity" bson:"quantity"`
}

type UserPreferences struct {
	UserID      string   `json:"user_id" bson:"user_id"`
	Cuisines    []string `json:"cuisines" bson:"cuisines"`
	MaxDistance float64  `json:"max_distance" bson:"max_distance"` // in kilometers
	MinRating   float64  `json:"min_rating" bson:"min_rating"`
	Location    Location `json:"location" bson:"location"`
}
