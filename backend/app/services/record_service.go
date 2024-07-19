package services

import (
	"sort"
	"word_bubble_popping/dto"
	"word_bubble_popping/models"
	"word_bubble_popping/repositories"
)

type IRecordService interface {
    GetRanking() (*[]models.Record, error)
    Create(createRecordInput dto.CreateRecordInput) (*models.Record, error)
}

type RecordService struct {
    repository repositories.IRecordRepository
}

func NewRecordService(repository repositories.IRecordRepository) IRecordService {
    return &RecordService{repository: repository}
}

func (s *RecordService) GetRanking() (*[]models.Record, error) {
    allRecords, err := s.repository.FindAll()
    if err != nil {
        return nil, err
    }

    sort.Slice(*allRecords, func(i, j int) bool {
        return (*allRecords)[i].Score > (*allRecords)[j].Score
    })

    if len(*allRecords) < 3 {
        return allRecords, nil
    }

    ranking := (*allRecords)[:3]

    return &ranking, nil
}

func (s *RecordService) Create(createRecordInput dto.CreateRecordInput) (*models.Record, error) {
    newRecord := models.Record{
        Name: createRecordInput.Name,
    }
    return s.repository.Create(newRecord)
}