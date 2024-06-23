package health

import (
	"net/http"

	healthservice "assignment/service/health"

	"github.com/labstack/echo/v4"
)

type Endpoint interface {
	Health(ctx echo.Context) error
}

type endpoint struct {
	service healthservice.Service
}

func NewEndpoint(healthService healthservice.Service) Endpoint {
	return &endpoint{service: healthService}
}

func (e *endpoint) Health(c echo.Context) error {
	request := new(HealthRequest)
	if err := c.Bind(request); err != nil {
		return err
	}

	status, err := e.service.Health(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, &HealthResponse{Status: "fail"})
	}

	return c.JSON(http.StatusOK, &HealthResponse{Status: status})
}
