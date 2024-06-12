package repository

import (
	"context"
	"errors"

	"github.com/IIGabriel/training-sheet-crud/internal/models"
	"github.com/IIGabriel/training-sheet-crud/internal/trainingSheet"
	mysqlConn "github.com/IIGabriel/training-sheet-crud/pkg/db/mysql"
	"gorm.io/gorm"
)

type trainingSheetRepository struct {
	mysql *gorm.DB
}

func NewTrainingSheetRepository() trainingSheet.Repository {
	return &trainingSheetRepository{mysql: mysqlConn.Mysql()}
}

func (t *trainingSheetRepository) Create(ctx context.Context, treino models.TrainingSheet) error {
	if err := t.mysql.WithContext(ctx).Create(&treino).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainingSheetRepository) Delete(ctx context.Context, id uint) error {
	if err := t.mysql.WithContext(ctx).Delete(&models.TrainingSheet{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (t *trainingSheetRepository) Get(ctx context.Context, id uint) (*models.TrainingSheet, error) {
	var treino models.TrainingSheet
	if err := t.mysql.WithContext(ctx).First(&treino, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &treino, nil
}

func (t *trainingSheetRepository) Update(ctx context.Context, treino models.TrainingSheet) error {
	if err := t.mysql.WithContext(ctx).Save(&treino).Error; err != nil {
		return err
	}
	return nil
}
