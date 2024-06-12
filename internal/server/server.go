package server

import (
	trainingSheetHttp "github.com/IIGabriel/training-sheet-crud/internal/trainingSheet/delivery/http"
	trainingSheetRepository "github.com/IIGabriel/training-sheet-crud/internal/trainingSheet/repository"
	trainingSheetUC "github.com/IIGabriel/training-sheet-crud/internal/trainingSheet/usecase"
	"github.com/labstack/echo/v4"
)

func (s *Server) MapControllers(app *echo.Echo) error {

	traningSheetRepo := trainingSheetRepository.NewTrainingSheetRepository()

	// UseCases
	stlUseCase := trainingSheetUC.NewTrainingSheetUseCase(traningSheetRepo)

	// Controllers
	stlController := trainingSheetHttp.NewTrainingSheetController(stlUseCase)

	// Routes
	trainingSheetHttp.MapTrainingSheetRoutes(app.Group("/training-sheet"), stlController)

	return nil
}
