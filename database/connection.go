package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"
    "github.com/rogerhendricks/go-vue/models"
)

func InitDB() *gorm.DB{
    db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }
    db.AutoMigrate(&models.User{}, &models.Device{}, &models.Lead{})
    return db
}