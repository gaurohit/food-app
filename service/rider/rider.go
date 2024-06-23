package rider

import (
	"assignment/model"
	"assignment/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RiderService interface {
	ResgisterRider(c echo.Context, request *model.Rider) *utils.ErrorHandler
	UpdateRiderLocation(c echo.Context, Location model.Location, riderId string) *utils.ErrorHandler
	GetRiderOrderHistory(c echo.Context, riderId string) ([]model.Order, *utils.ErrorHandler)
	NearestRider(c echo.Context, resturantId string) (*model.Rider, *utils.ErrorHandler)
}

type riderService struct {
	config      *viper.Viper
	mongoClient *mongo.Client
}

func NewRiderService(config *viper.Viper, mongoClient *mongo.Client) RiderService {
	return &riderService{config: config, mongoClient: mongoClient}
}

func (r *riderService) ResgisterRider(c echo.Context, request *model.Rider) *utils.ErrorHandler {
	request.ID = uuid.NewString()
	collection := r.mongoClient.Database("test").Collection("riders")

	err := model.Register(c, request, *collection)
	if err != nil {
		return &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	return nil

}

func (r *riderService) UpdateRiderLocation(c echo.Context, Location model.Location, riderId string) *utils.ErrorHandler {
	filter := bson.M{"_id": riderId}

	collection := r.mongoClient.Database("test").Collection("riders")

	err := model.UpdateRiderLocation(c, Location, *collection, filter)

	if err != nil {
		return &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	return nil
}

func (r *riderService) GetRiderOrderHistory(c echo.Context, riderId string) ([]model.Order, *utils.ErrorHandler) {

	filter := bson.M{
		"rider_id": riderId,
	}

	collection := r.mongoClient.Database("test").Collection("orders")

	response, err := model.GetRiderOrderHistory(c, filter, *collection)
	if err != nil {
		return nil, &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage, Code: err.Code}
	}

	return *response, nil
}

func (r *riderService) NearestRider(c echo.Context, resturantId string) (*model.Rider, *utils.ErrorHandler) {

	collection := r.mongoClient.Database("test").Collection("restaurants")

	restaurant, err := model.GetRestaurant(c, resturantId, *collection)

	if err != nil {
		return nil, &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage, Code: err.Code}
	}

	Collection := r.mongoClient.Database("test").Collection("riders")

	rider, errRider := model.NearestRider(c, restaurant.Location, Collection)

	if errRider != nil {
		return nil, &utils.ErrorHandler{Message: errRider.Message, DevMessage: errRider.DevMessage, Code: errRider.Code}
	}

	return rider, nil

}
