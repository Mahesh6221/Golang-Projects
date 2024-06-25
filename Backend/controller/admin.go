package controllers

import (
	"BACKEND/database"
	"BACKEND/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetEmployeeAsset(c *gin.Context) {
    // Get the empID parameter from the URL path
    empID := c.Param("empID")

    if empID == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "empID parameter is required"})
        return
    }

    client := database.InitializeMongoDB()
    EmployeeDetailsCollection := database.GetCollection(client, "EmployeeDetails")
    AssetCollection := database.GetCollection(client, "AssetDetails")

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    // Fetching employee details by empID
    var employeeDetails model.User
    if err := EmployeeDetailsCollection.FindOne(ctx, bson.M{"empID": empID}).Decode(&employeeDetails); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving employee details"})
        return
    }

    // Fetching asset details by empID
    var assetDetails model.AssetDetail
    if err := AssetCollection.FindOne(ctx, bson.M{"empID": empID}).Decode(&assetDetails); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Error retrieving asset details"})
        return
    }

    // Combine employee and asset details into a single response
    data := map[string]interface{}{
        "details": employeeDetails,
        "assets":  assetDetails,
    }

    c.JSON(http.StatusOK, gin.H{"data": data})
}
