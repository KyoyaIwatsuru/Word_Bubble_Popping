package models

import "gorm.io/gorm"

type Word struct {
    gorm.Model
    WordEn string `gorm:"not null"`
    WordJa string `gorm:"not null"`
}