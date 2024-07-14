package repositories

import (
	"word_bubble_popping/models"

	"gorm.io/gorm"
)

type IWordRepository interface {
    FindAll() (*[]models.Word, error)
}

type WordMemoryRepository struct {
    words []models.Word
}

func NewWordMemoryRepository(words []models.Word) IWordRepository {
    return &WordMemoryRepository{words: words}
}

func (r *WordMemoryRepository) FindAll() (*[]models.Word, error) {
    return &r.words, nil
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