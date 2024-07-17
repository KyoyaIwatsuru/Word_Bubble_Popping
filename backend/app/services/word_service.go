package services

import (
	"errors"
	"math/rand"
	"time"
	"word_bubble_popping/models"
	"word_bubble_popping/repositories"
)

type IWordService interface {
    GetRandom() (*[]models.Word, error)
}

type WordService struct {
    repository repositories.IWordRepository
}

func NewWordService(repository repositories.IWordRepository) IWordService {
    return &WordService{repository: repository}
}

func (s *WordService) GetRandom() (*[]models.Word, error) {
    allWords, err := s.repository.FindAll()
    if err != nil {
        return nil, err
    }

    rand.New(rand.NewSource(time.Now().UnixNano()))
    rand.Shuffle(len(*allWords), func(i, j int) {
        (*allWords)[i], (*allWords)[j] = (*allWords)[j], (*allWords)[i]
    })

    if len(*allWords) < 3 {
        return nil, errors.New("Not enough words")
    }

    randWords := (*allWords)[:3]

    return &randWords, nil
}