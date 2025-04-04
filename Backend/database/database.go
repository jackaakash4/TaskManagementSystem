package database

import (
    "fmt"
    "log"

    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

//Global DB variable
var DB *gorm.DB

//connect to database

func ConnectDatabase() {
    dsn := "host=localhost user=postgres password=123456789 dbname=TaskManagementSystemDB port=8080 sslmode=disable"

    var err error
    
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

    if err != nil {
        log.Fatal("Failed to connect to database", err)
    }

    fmt.Println("Database connected sucessfully")
}

