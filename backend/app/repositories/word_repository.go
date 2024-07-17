package repositories

import (
	"word_bubble_popping/models"

	"gorm.io/gorm"
)

type IWordRepository interface {
    FindAll() (*[]models.Word, error)
}

type WordRepository struct {
    db *gorm.DB
}

func NewWordRepository(db *gorm.DB) IWordRepository {
    return &WordRepository{db: db}
}

func (r *WordRepository) FindAll() (*[]models.Word, error) {
    var words []models.Word
    result := r.db.Find(&words)
    if result.Error != nil {
        return nil, result.Error
    }
    return &words, nil
}