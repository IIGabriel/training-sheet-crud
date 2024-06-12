package trainingSheet

import (
	"context"
	"github.com/IIGabriel/training-sheet-crud/internal/models"
)

type UseCase interface {
	Create(context.Context, models.TrainingSheet) error
	Delete(ctx context.Context, id uint) error
	Get(ctx context.Context, id uint) (*models.TrainingSheet, error)
	Update(context.Context, models.TrainingSheet) error
}
