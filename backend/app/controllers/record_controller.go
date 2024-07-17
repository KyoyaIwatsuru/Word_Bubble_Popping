package controllers

import (
	"net/http"
	"word_bubble_popping/dto"
	"word_bubble_popping/services"

	"github.com/gin-gonic/gin"
)

type IRecordController interface {
    GetRanking(ctx *gin.Context)
    Create(ctx *gin.Context)
}

type RecordController struct {
    service services.IRecordService
}

func NewRecordController(service services.IRecordService) IRecordController {
    return &RecordController{service: service}
}

func (c *RecordController) GetRanking(ctx *gin.Context) {
    ranking, err := c.service.GetRanking()
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": ranking})
}

func (c *RecordController) Create(ctx *gin.Context) {
    var input dto.CreateRecordInput
    if err := ctx.ShouldBindJSON(&input); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    newRecord, err := c.service.Create(input)
    if err != nil {
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }

    ctx.JSON(http.StatusCreated, gin.H{"data": newRecord})
}