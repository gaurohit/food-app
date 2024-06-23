package model

import (
	"assignment/utils"
	"context"

	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Register(c echo.Context, request interface{}, collection mongo.Collection) *utils.ErrorHandler {

	_, err := collection.InsertOne(c.Request().Context(), &request)
	if err != nil {
		return &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(),Code:400}
	}

	return nil

}

func UpdateRiderLocation(c echo.Context, location Location, collection mongo.Collection, filter primitive.M) *utils.ErrorHandler {

	update := bson.M{"$set": bson.M{"location": location}}

	result := collection.FindOneAndUpdate(c.Request().Context(), filter, update)
	if result.Err() != nil {
		if result.Err() == mongo.ErrNoDocuments {
			return &utils.ErrorHandler{Message: utils.DATA_NOT_FOUND, DevMessage: result.Err().Error(),Code:404}
		}
		return &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: result.Err().Error(),Code:400}
	}

	return nil
}

func GetRiderOrderHistory(c echo.Context, filter primitive.M, collection mongo.Collection) (*[]Order, *utils.ErrorHandler) {
	result := new([]Order)
	// need to pagination

	cursor, err := collection.Find(c.Request().Context(), filter)
	if err != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}
	defer cursor.Close(c.Request().Context())

	errCursor := cursor.All(c.Request().Context(), result)
	if errCursor != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: errCursor.Error(), Code: 400}
	}

	if len(*result) == 0 {
		return nil, &utils.ErrorHandler{Message: utils.DATA_NOT_FOUND, Code: 404}
	}

	return result, nil

}

func SuggestRestaurant(c echo.Context, preferences UserPreferences, collection mongo.Collection) ([]Restaurant, *utils.ErrorHandler) {
	var restaurants []Restaurant
	// need to pagination

	filter := bson.M{
		"cuisine": bson.M{"$in": preferences.Cuisines},
		"rating":  bson.M{"$gte": preferences.MinRating},
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{preferences.Location.Longitude, preferences.Location.Latitude},
				},
				"$maxDistance": preferences.MaxDistance,
			},
		},
	}
	cursor, err := collection.Find(c.Request().Context(), filter, options.Find().SetLimit(50))
	if err != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}
	defer cursor.Close(context.TODO())

	if err := cursor.All(c.Request().Context(), &restaurants); err != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}

	if len(*restaurants) == 0 {
		return nil, &utils.ErrorHandler{Message: utils.DATA_NOT_FOUND, Code: 404}
	}

	return restaurants, nil
}

func GetRestaurant(c echo.Context, restaurantId string, collection mongo.Collection) (*Restaurant, *utils.ErrorHandler) {

	filter := bson.M{"_id": restaurantId}

	restaurant :=  new(Restaurant)
	err := collection.FindOne(c.Request().Context(), filter).Decode(&restaurant)
	if err != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}

	if restaurant == nil{
		return nil, &utils.ErrorHandler{Message: utils.DATA_NOT_FOUND, Code: 404}
	}

	return restaurant, nil
}

func GetUser(c echo.Context, filter primitive.M, Client mongo.Client) (*User, *utils.ErrorHandler) {

	user := new(User)

	collection := Client.Database("test").Collection("users")

	err := collection.FindOne(c.Request().Context(), filter).Decode(user)
	if err != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}

	if user == nil{
		return nil, &utils.ErrorHandler{Message: utils.DATA_NOT_FOUND, Code: 404}
	}

	return user, nil
}

func AcceptOrder(c echo.Context, orderDetails Order, collection mongo.Collection) *utils.ErrorHandler {

		_, err := collection.InsertOne(c.Request().Context(), orderDetails)
	if err != nil {
		return &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}

	return nil

}

func GetUserOrderHistory(c echo.Context, filter primitive.M, collection mongo.Collection) ([]Order, *utils.ErrorHandler) {
	result := new([]Order)
	// need to pagination

	cursor, err := collection.Find(c.Request().Context(), filter, options.Find().SetLimit(10))
	if err != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}
	defer cursor.Close(c.Request().Context())

	errCursor := cursor.All(c.Request().Context(), result)
	if errCursor != nil {
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: errCursor.Error(), Code: 400}
	}

	if len(*result) == 0 {
		return nil, &utils.ErrorHandler{Message: utils.DATA_NOT_FOUND, Code: 404}
	}

	return *result, nil

}

func NearestRider(c echo.Context, points Location, collection *mongo.Collection) (*Rider, *utils.ErrorHandler) {
	rider := new(Rider)
	filter := bson.M{
		"location": bson.M{
			"$nearSphere": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{points.Longitude, points.Latitude},
				},
				"$maxDistance": 5000, // assuming distance within 5 KM
			},
		},
	}

	err := collection.FindOne(c.Request().Context(), filter, options.FindOne().SetSort(bson.M{
		"location": bson.M{
			"$meta": "textScore",
		},
	})).Decode(rider)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 404}
		}
		return nil, &utils.ErrorHandler{Message: utils.SOMETHING_WENT_WRONG, DevMessage: err.Error(), Code: 400}
	}

	return rider, nil

}
