package health

import (
	"github.com/labstack/echo/v4"
)

type Service interface {
	Health(c echo.Context) (string, error)
}

type service struct {
}

func (s *service) Health(c echo.Context) (string, error) {
	return "pass", nil
}
func NewService() Service {
	return &service{}
}
