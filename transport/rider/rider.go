package rider

import (
	"assignment/model"
	riderService "assignment/service/rider"
	"assignment/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type riderEndpoint struct {
	riderService riderService.RiderService
}

type RiderEndpoint interface {
	RegisterRider(c echo.Context) error
	UpdateRiderLocation(c echo.Context) error
	GetRiderOrderHistory(c echo.Context) error
	NearestRider(c echo.Context) error
}

func NewRiderEndpoint(riderService riderService.RiderService) RiderEndpoint {
	return &riderEndpoint{riderService: riderService}
}

func (r *riderEndpoint) RegisterRider(c echo.Context) error {
	rider := new(model.Rider)
	if err := c.Bind(rider); err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: err.Error()})
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: utils.INVALID_REQUEST})
	}

	validationError := utils.Validator(rider)
	if validationError != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: validationError.Message})
		return c.JSON(http.StatusBadRequest, &utils.GenericResponse{Message: validationError.Message})
	}

	errRegister := r.riderService.ResgisterRider(c, rider)
	if errRegister != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: errRegister.Message})
		return c.JSON(errRegister.Code, utils.GenericResponse{Message: errRegister.Message})
	}

	return c.JSON(http.StatusCreated, utils.GenericResponse{Message: "Rider registered successfully"})

}

func (r *riderEndpoint) UpdateRiderLocation(c echo.Context) error {
	updateLoaction := new(model.Location)
	riderId := c.Param("id")
	if err := c.Bind(updateLoaction); err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: err.Error()})
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: utils.DATA_NOT_FOUND})
	}

	validationError := utils.Validator(updateLoaction)
	if validationError != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: validationError.Message})
		return c.JSON(http.StatusBadRequest, &utils.GenericResponse{Message: validationError.Message})
	}

	err := r.riderService.UpdateRiderLocation(c, *updateLoaction, riderId)
	if err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: err.Message})
		return c.JSON(err.Code, utils.GenericResponse{Message: err.Message})
	}

	return c.JSON(http.StatusOK, utils.GenericResponse{Message: "Rider registered successfully"})

}

func (r *riderEndpoint) GetRiderOrderHistory(c echo.Context) error {
	riderId := c.Param("id")
	response, err := r.riderService.GetRiderOrderHistory(c, riderId)
	if err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: err.Message})
		return c.JSON(err.Code, utils.GenericResponse{Message: err.Message})
	}

	return c.JSON(http.StatusOK, response)
}

func (r *riderEndpoint) NearestRider(c echo.Context) error {
	restaurantId := c.Param("id")
	response, err := r.riderService.NearestRider(c, restaurantId)
	if err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: errRegister.Message})
		return c.JSON(err.Code, utils.GenericResponse{Message: err.Message})
	}

	return c.JSON(http.StatusOK, response)
}
