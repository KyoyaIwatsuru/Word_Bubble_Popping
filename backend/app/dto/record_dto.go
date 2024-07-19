package dto

type CreateRecordInput struct {
    Name string `json:"name" binding:"required,min=2"`
}