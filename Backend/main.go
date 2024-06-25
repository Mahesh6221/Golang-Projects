package main

import (
	controller "BACKEND/controller"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Port := os.Getenv("PORT")
	if Port == "" {
		Port = "8080"
	}
	r := gin.Default()
	r.POST("/login", controller.Login())
	r.GET("/users", controller.GetAllUsers)
	r.GET("/admin/:empID", controller.GetEmployeeAsset)
	r.POST("/upload",controller.Uploadfile)
	r.GET("/asset-details", controller.GetAssetDetails)
	r.Run(":" + Port)
}