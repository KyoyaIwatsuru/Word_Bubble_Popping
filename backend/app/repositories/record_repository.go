package repositories

import (
	"word_bubble_popping/models"

	"gorm.io/gorm"
)

type IRecordRepository interface {
    FindAll() (*[]models.Record, error)
    Create(newRecord models.Record) (*models.Record, error)
}

type RecordRepository struct {
    db *gorm.DB
}

func NewRecordRepository(db *gorm.DB) IRecordRepository {
    return &RecordRepository{db: db}
}

func (r *RecordRepository) FindAll() (*[]models.Record, error) {
    var records []models.Record
    result := r.db.Find(&records)
    if result.Error != nil {
        return nil, result.Error
    }
    return &records, nil
}

func (r *RecordRepository) Create(newRecord models.Record) (*models.Record, error) {
    result := r.db.Create(&newRecord)
    if result.Error != nil {
        return nil, result.Error
    }
    return &newRecord, nil
}