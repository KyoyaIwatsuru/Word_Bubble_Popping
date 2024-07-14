package main

import (
	"word_bubble_popping/controllers"
	"word_bubble_popping/infra"
	// "word_bubble_popping/models"
	"word_bubble_popping/repositories"
	"word_bubble_popping/services"

	"github.com/gin-gonic/gin"
)

func main() {
    infra.Initialize()
    db := infra.SetupDB()

    // words := []models.Word{
    //     {WordId: 1, WordEn: "apple", WordJa: "りんご"},
    //     {WordId: 2, WordEn: "banana", WordJa: "バナナ"},
    //     {WordId: 3, WordEn: "cherry", WordJa: "さくらんぼ"},
    //     {WordId: 4, WordEn: "orange", WordJa: "オレンジ"},
    //     {WordId: 5, WordEn: "strawberry", WordJa: "いちご"},
    //     {WordId: 6, WordEn: "grape", WordJa: "ぶどう"},
    //     {WordId: 7, WordEn: "melon", WordJa: "メロン"},
    //     {WordId: 8, WordEn: "peach", WordJa: "もも"},
    //     {WordId: 9, WordEn: "pear", WordJa: "なし"},
    //     {WordId: 10, WordEn: "watermelon", WordJa: "すいか"},
    // }

    // wordRepository := repositories.NewWordMemoryRepository(words)
    wordRepository := repositories.NewWordRepository(db)
    wordService := services.NewWordService(wordRepository)
    wordController := controllers.NewWordController(wordService)

    r := gin.Default()
    r.GET("/words", wordController.FindRandom)
    r.Run(":8080")
}