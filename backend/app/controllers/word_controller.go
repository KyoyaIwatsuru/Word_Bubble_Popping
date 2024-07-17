package controllers

import (
	"net/http"
	"word_bubble_popping/services"

	"github.com/gin-gonic/gin"
)

type IWordController interface {
    GetRandom(ctx *gin.Context)
}

type WordController struct {
    service services.IWordService
}

func NewWordController(service services.IWordService) IWordController {
    return &WordController{service: service}
}

func (c *WordController) GetRandom(ctx *gin.Context) {
    words, err := c.service.GetRandom()
    if err != nil {
        if err.Error() == "Not enough words" {
            ctx.JSON(http.StatusNotImplemented, gin.H{"error": err.Error()})
            return
        }
        ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
        return
    }

    ctx.JSON(http.StatusOK, gin.H{"data": words})
}