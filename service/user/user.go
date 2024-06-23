package user

import (
	"assignment/model"
	"assignment/utils"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService interface {
	Resgister(c echo.Context, request *model.User) *utils.ErrorHandler
	GetUserOrderHistory(c echo.Context, userId string) ([]model.Order, *utils.ErrorHandler)
}

type userService struct {
	config      *viper.Viper
	mongoClient *mongo.Client
}

func NewUserService(config *viper.Viper, mongoClient *mongo.Client) UserService {
	return &userService{config: config, mongoClient: mongoClient}
}

func (u *userService) Resgister(c echo.Context, request *model.User) *utils.ErrorHandler {
	request.ID = uuid.NewString()
	
	collection := u.mongoClient.Database("test").Collection("users")
	
	err := model.Register(c, request, *collection)
	if err != nil {
		return &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage}
	}

	return nil

}

func (r *userService) GetUserOrderHistory(c echo.Context, userId string) ([]model.Order, *utils.ErrorHandler) {
	filter := bson.M{"_id": userId}
	collection := r.mongoClient.Database("test").Collection("orders")
	response, err := model.GetUserOrderHistory(c, filter, *collection)
	if err != nil {
		return nil, &utils.ErrorHandler{Message: err.Message, DevMessage: err.DevMessage, Code: err.Code}
	}
	return response, nil
}
