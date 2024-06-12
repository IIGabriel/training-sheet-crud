package trainingSheet

import (
	"github.com/labstack/echo/v4"
)

type Controller interface {
	Create(ctx echo.Context) error
	Update(ctx echo.Context) error
	Get(ctx echo.Context) error
	Delete(ctx echo.Context) error
}
