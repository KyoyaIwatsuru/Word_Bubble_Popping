package models

import "gorm.io/gorm"

type Record struct {
    gorm.Model
    Name        string  `gorm:"not null;unique"`
    Score       int     `gorm:"not null;default:0"`
    words       []Word
}