package usecase

import (
	"context"
	"github.com/IIGabriel/training-sheet-crud/internal/models"
	"github.com/IIGabriel/training-sheet-crud/internal/trainingSheet"
)

type trainingSheetUseCase struct {
	repo trainingSheet.Repository
}

func NewTrainingSheetUseCase(repo trainingSheet.Repository) trainingSheet.UseCase {
	return &trainingSheetUseCase{repo}
}

func (t *trainingSheetUseCase) Create(ctx context.Context, treino models.TrainingSheet) error {
	return t.repo.Create(ctx, treino)
}

func (t *trainingSheetUseCase) Delete(ctx context.Context, id uint) error {
	return t.repo.Delete(ctx, id)
}

func (t *trainingSheetUseCase) Get(ctx context.Context, id uint) (*models.TrainingSheet, error) {
	return t.repo.Get(ctx, id)
}

func (t *trainingSheetUseCase) Update(ctx context.Context, treino models.TrainingSheet) error {
	return t.repo.Update(ctx, treino)
}
