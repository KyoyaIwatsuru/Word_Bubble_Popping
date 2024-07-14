package main

import (
	"log"
	"word_bubble_popping/infra"
	"word_bubble_popping/models"

	"gorm.io/gorm"
)

func main() {
    infra.Initialize()
    db := infra.SetupDB()

    if err := db.AutoMigrate(&models.Word{}); err != nil {
        panic("Failed to migrate database")
    }

    seedData(db)
}

func seedData(db *gorm.DB) {
    var count int64
    db.Model(&models.Word{}).Count(&count)

    if count < 10 {
        words := []models.Word{
            {WordEn: "apple", WordJa: "りんご"},
            {WordEn: "banana", WordJa: "バナナ"},
            {WordEn: "cherry", WordJa: "さくらんぼ"},
            {WordEn: "orange", WordJa: "オレンジ"},
            {WordEn: "strawberry", WordJa: "いちご"},
            {WordEn: "grape", WordJa: "ぶどう"},
            {WordEn: "melon", WordJa: "メロン"},
            {WordEn: "peach", WordJa: "もも"},
            {WordEn: "pear", WordJa: "なし"},
            {WordEn: "watermelon", WordJa: "すいか"},
        }

        for _, word := range words {
            result := db.Create(&word)
            if result.Error != nil {
                panic("Failed to seed data")
            }
        }
    } else {
        log.Println("Initial data already exists. Skipping seed data insertion.")
    }
}