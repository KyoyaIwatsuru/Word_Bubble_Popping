package repositories

import "word_bubble_popping/models"

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