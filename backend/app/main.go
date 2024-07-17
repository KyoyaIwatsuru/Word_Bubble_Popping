package main

import (
	"word_bubble_popping/controllers"
	"word_bubble_popping/infra"
	"word_bubble_popping/repositories"
	"word_bubble_popping/services"

	"github.com/gin-gonic/gin"
)

func main() {
    infra.Initialize()
    db := infra.SetupDB()

    wordRepository := repositories.NewWordRepository(db)
    wordService := services.NewWordService(wordRepository)
    wordController := controllers.NewWordController(wordService)

    recordRepository := repositories.NewRecordRepository(db)
    recordService := services.NewRecordService(recordRepository)
    recordController := controllers.NewRecordController(recordService)

    r := gin.Default()
    wordRouter := r.Group("/words")
    recordRouter := r.Group("/records")

    wordRouter.GET("", wordController.GetRandom)
    recordRouter.GET("", recordController.GetRanking)
    recordRouter.POST("", recordController.Create)
    r.Run(":8080")
}