package user

import (
	"assignment/model"
	userService "assignment/service/user"
	"assignment/utils"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type userEndpoint struct {
	userService userService.UserService
}

type UserEnpoint interface {
	Register(c echo.Context) error
	GetUserOrderHistory(c echo.Context) error
}

func NewUserEndpoint(userService userService.UserService) UserEnpoint {
	return &userEndpoint{userService: userService}
}

func (u *userEndpoint) Register(c echo.Context) error {
	user := new(model.User)

	err := c.Bind(user)
	if err != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: err.Error()})
		return c.JSON(http.StatusBadRequest, utils.GenericResponse{Message: utils.INVALID_REQUEST})
	}

	validationError := utils.Validator(user)
	if validationError != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: validationError.Message})
		return c.JSON(http.StatusBadRequest, &utils.GenericResponse{Message: validationError.Message})
	}

	errRegister := u.userService.Resgister(c, user)
	if errRegister != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: errRegister.Message})
		return c.JSON(errRegister.Code, utils.GenericResponse{Message: errRegister.Message})
	}

	return c.JSON(http.StatusCreated, utils.GenericResponse{Message: "user registered successfully"})

}

func (r *userEndpoint) GetUserOrderHistory(c echo.Context) error {
	riderId := c.Param("id")
	response, errRegister := r.userService.GetUserOrderHistory(c, riderId)
	if errRegister != nil {
		log.Println(c.Request().RequestURI, &utils.GenericResponse{Message: errRegister.Message})
		return c.JSON(errRegister.Code, utils.GenericResponse{Message: errRegister.Message})
	}

	return c.JSON(http.StatusOK, response)
}
