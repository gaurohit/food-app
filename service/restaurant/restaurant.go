package restaurant

import (
	"assignment/model"
	"assignment/requestresponse"
	"assignment/utils"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"

	"go.mongodb.org/mongo-driver/mongo"
)

type RestaurantService interface {
	Resgister(c echo.Context, request *model.Restaurant) *utils.ErrorHandler
	SuggestRestaurant(c echo.Context, request model.UserPreferences) ([]model.Restaurant, *utils.ErrorHandler)
	GetRestaurantMenu(c echo.Context, restaurantId string) (*model.Restaurant, *utils.ErrorHandler)
	AcceptOrder(c echo.Context, orderDetails requestresponse.Order) *utils.ErrorHandler
}

type restaurantService struct {
	config      *viper.Viper
	mongoClient *mongo.Client
}

func NewRestaurantService(config *viper.Viper, mongoClient *mongo.Client) RestaurantService {
	return &restaurantService{config: config, mongoClient: mongoClient}
}

func (u *restaurantService) Resgister(c echo.Context, request *model.Restaurant) *utils.ErrorHandler {
	request.ID = uuid.NewString()

	collection := u.mongoClient.Database("test").Collection("restaurant")
	err := model.Register(c, request, *collection)
	if err != nil {
		return &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	return nil

}

func (u *restaurantService) SuggestRestaurant(c echo.Context, request model.UserPreferences) ([]model.Restaurant, *utils.ErrorHandler) {

	collection := u.mongoClient.Database("test").Collection("restaurants")

	response, err := model.SuggestRestaurant(c, request, *collection)
	if err != nil {
		return nil, &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	return response, nil

}

func (u *restaurantService) GetRestaurantMenu(c echo.Context, restaurantId string) (*model.Restaurant, *utils.ErrorHandler) {
	collection := u.mongoClient.Database("test").Collection("restaurants")

	response, err := model.GetRestaurant(c, restaurantId, *collection)

	if err != nil {
		return nil, &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	return response, nil
}

func (u *restaurantService) AcceptOrder(c echo.Context, orderDetails requestresponse.Order) *utils.ErrorHandler {

	filter := bson.M{"_id": orderDetails.UserID}
	
	user, err := model.GetUser(c, filter, *u.mongoClient)

	if err != nil {
		return &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	collection := u.mongoClient.Database("test").Collection("restaurants")

	restaurant, err := model.GetRestaurant(c, orderDetails.RestaurantID, *collection)

	prices := make(map[string]float64)

	for _, v := range restaurant.Menu {
		prices[v.ID] = v.Price
	}

	if err != nil {
		return &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	order := new(model.Order)

	totalPrices := 0.00

	for _, value := range orderDetails.Items {
		totalPrices += prices[value.MenuItemID]
	}

	order.ID = uuid.NewString()
	order.Items = orderDetails.Items
	order.UserID = user.ID
	order.Status = "Pending"
	order.TotalPrice = totalPrices
	order.Createdat = time.Now()
	order.Updatedat = time.Now()

	collection = u.mongoClient.Database("test").Collection("orders")

	err = model.AcceptOrder(c, *order, *collection)

	if err != nil {
		return &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	return nil
}
