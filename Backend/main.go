package main

import (

    "github.com/jackaakash4/TaskManagementSystem/database"
    
    "github.com/gin-gonic/gin"
)

func main() {
    
    //connect to database
    
    database.ConnectDatabase()

    //Initialize router
    
    r := gin.Default()

    r.GET("/", func(c *gin.Context){
        c.JSON(200, gin.H{
            "message": "Helllo world",
        })
    })

    r.Run(":443")

    

}

