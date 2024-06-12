package http

import (
	"github.com/IIGabriel/training-sheet-crud/internal/models"
	"github.com/IIGabriel/training-sheet-crud/internal/trainingSheet"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type trainingSheetController struct {
	UseCase trainingSheet.UseCase
}

func NewTrainingSheetController(uc trainingSheet.UseCase) trainingSheet.Controller {
	return &trainingSheetController{uc}
}

func (t *trainingSheetController) Create(ctx echo.Context) error {
	var sheet models.TrainingSheet
	if err := ctx.Bind(&sheet); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := t.UseCase.Create(ctx.Request().Context(), sheet); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not create training sheet"})
	}
	return ctx.JSON(http.StatusCreated, sheet)
}

func (t *trainingSheetController) Update(ctx echo.Context) error {
	var sheet models.TrainingSheet
	if err := ctx.Bind(&sheet); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request"})
	}
	if err := t.UseCase.Update(ctx.Request().Context(), sheet); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not update training sheet"})
	}
	return ctx.JSON(http.StatusOK, sheet)
}

func (t *trainingSheetController) Get(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	sheet, err := t.UseCase.Get(ctx.Request().Context(), uint(id))
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not get training sheet"})
	}
	return ctx.JSON(http.StatusOK, sheet)
}

func (t *trainingSheetController) Delete(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	if err := t.UseCase.Delete(ctx.Request().Context(), uint(id)); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": "could not delete training sheet"})
	}
	return ctx.NoContent(http.StatusNoContent)
}
