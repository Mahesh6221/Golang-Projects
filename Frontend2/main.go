package main

import (
	"FRONTEND2/handler"
	"FRONTEND2/middleware"
	"fmt"
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
	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/assets", "./assets/")
	if Port == "" {
		fmt.Println("unable to get Port Number")
		// here i am using same port
		Port = "9090"
	}
	r.Use(middleware.CorsMiddleware())
	r.GET("/login", handler.Login)
	r.POST("/login", handler.PostLogin)
	r.GET("/logout", handler.UserLogout)
	r.GET("/admin", handler.CheckTokenExpiration(), handler.GetAdmin)
	r.GET("/users", handler.CheckTokenExpiration(), handler.GetUsers)
	r.Run(":" + Port)
}
