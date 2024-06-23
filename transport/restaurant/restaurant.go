package restaurant

import (
	"assignment/model"
	"assignment/requestresponse"
	restaurantService "assignment/service/restaurant"
	"assignment/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type restaurantEndpoint struct {
	restaurantService restaurantService.RestaurantService
}

type RestaurantEnpoint interface {
	Register(c echo.Context) error
	SuggestRestaurant(c echo.Context) error
	GetRestaurantMenu(c echo.Context) error
	AcceptOrder(c echo.Context) error
}

func NewRestaurantEndpoint(restaurantService restaurantService.RestaurantService) RestaurantEnpoint {
	return &restaurantEndpoint{restaurantService: restaurantService}
}

func (r *restaurantEndpoint) Register(c echo.Context) error {
	restaurant := new(model.Restaurant)

	err := c.Bind(restaurant)
	if err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: err.Error()})
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: utils.INVALID_REQUEST})
	}

	validationError := utils.Validator(restaurant)
	if validationError != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: validationError.Message})
		return c.JSON(http.StatusBadRequest, &utils.GenericResponse{Message: validationError.Message})
	}

	errRegister := r.restaurantService.Resgister(c, restaurant)
	if errRegister != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: errRegister.Message})
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: errRegister.Message})
	}

	return c.JSON(http.StatusCreated, utils.GenericResponse{Message: "user registered successfully"})

}

func (u *restaurantEndpoint) SuggestRestaurant(c echo.Context) error {
	preferences := new(model.UserPreferences)

	err := c.Bind(preferences)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: "data not valid"})
	}

	response, errSuggest := u.restaurantService.SuggestRestaurant(c, *preferences)
	if errSuggest != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: errSuggest.Message})
		return c.JSON(errSuggest.Code, utils.GenericResponse{Message: errSuggest.Message})
	}

	return c.JSON(http.StatusCreated, response)

}

func (r *restaurantEndpoint) GetRestaurantMenu(c echo.Context) error {
	restaurantId := c.Param("id")

	response, err := r.restaurantService.GetRestaurantMenu(c, restaurantId)
	if err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: err.Message})
		return c.JSON(err.Code, utils.GenericResponse{Message: err.Message})
	}

	return c.JSON(http.StatusOK, response.Menu)
}

func (r *restaurantEndpoint) AcceptOrder(c echo.Context) error {

	order := new(requestresponse.Order)
	err := c.Bind(order)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: utils.INVALID_REQUEST})
	}

	errAcceptOrder := r.restaurantService.AcceptOrder(c, *order)
	if errAcceptOrder != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: errAcceptOrder.Message})
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: errAcceptOrder.Message})
	}

	return c.JSON(http.StatusCreated, utils.GenericResponse{Message: "order accpeted successfully"})
}
