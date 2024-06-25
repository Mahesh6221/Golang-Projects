package controllers

import (
	"BACKEND/database"
	helper "BACKEND/helper"
	"BACKEND/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var foundUser model.User

func Login() gin.HandlerFunc {

	return func(c *gin.Context) {
		client := database.InitializeMongoDB()
		UserCollection := database.GetCollection(client, "EmployeeDetails")
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user model.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		err := UserCollection.FindOne(ctx, bson.M{"mailID": user.MailID}).Decode(&foundUser)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Entered email wrong"})
				return
			}
			c.JSON(http.StatusInternalServerError, gin.H{"error": "An error occurred"})
			return
		}

		if foundUser.Password != user.Password {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Please Enter a Valid password"})
			return
		}
		foundUser.Token, foundUser.Refresh_token, _ = helper.GenerateAllTokens(foundUser.MailID, foundUser.Password, foundUser.ID)
		c.JSON(http.StatusOK, foundUser)
	}
}
