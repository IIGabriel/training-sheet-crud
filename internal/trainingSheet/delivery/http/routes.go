package http

import (
	"github.com/IIGabriel/training-sheet-crud/internal/trainingSheet"
	"github.com/labstack/echo/v4"
)

func MapTrainingSheetRoutes(group *echo.Group, h trainingSheet.Controller) {
	group.POST("", h.Create)
	group.PUT("/:id", h.Update)
	group.GET("/:id", h.Get)
	group.DELETE("/:id", h.Delete)
}
