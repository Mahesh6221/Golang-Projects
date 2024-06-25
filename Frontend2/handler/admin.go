package handler

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAdmin(c *gin.Context) {
	empID := c.Query("empID")
	if empID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Employee ID is required"})
		return
	}

	// Make HTTP GET request to fetch employee and asset details
	url := "http://localhost:8080/admin/:empID" + empID
	employeeAssetsRes, err := http.Get(url)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch employee and asset details"})
		return
	}
	defer employeeAssetsRes.Body.Close()

	body, err := io.ReadAll(employeeAssetsRes.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to read response body"})
		return
	}

	// Unmarshal the response body into the struct
	var employeeAssets map[string]interface{}
	err = json.Unmarshal(body, &employeeAssets)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse response body"})
		return
	}

	// Respond with the fetched employee and asset details
	c.HTML(http.StatusOK, "admin.html", gin.H{"Data": employeeAssets})
}
