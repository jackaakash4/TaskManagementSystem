package models

import "gorm.io/gorm"

type User struct {
    gorm.Model
    ID          uint `json:"primarykey"`
    Name        string `json:"not null"`
    Email       string `json:"unique;not null"`
    Password    string `json:"not null"`
    Role        string `json:"default:'Viewer'"`
}


