package handler

import (
	"FRONTEND2/model"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// func GetUsers(c *gin.Context) {
// 	// fmt.Println("result data", result)

// 	if result.Type == "User" {
// 		// Assuming result contains the EmployeeID
// 		employeeID := result.EmpID
// 		url := fmt.Sprintf("http://localhost:8080/admin/%s", employeeID)
// 		// fmt.Println("employeeID:", employeeID)
// 		res1, err := http.Get(url)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		// fmt.Printf("res1: %v\n", res1.Body)
// 		defer res1.Body.Close()

// 		var responseData1 map[string]interface{}
// 		err = json.NewDecoder(res1.Body).Decode(&responseData1)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Extract asset details from the response
// 		assetDetails1, ok := responseData1["data"].(map[string]interface{})
// 		if !ok {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data format error"})
// 			return
// 		}
// 		// fmt.Printf("assetDetails1: %v\n", assetDetails1)
// 		// fmt.Printf("assetDetails1: %v\n", assetDetails1)
// 		c.HTML(http.StatusOK, "user.html", gin.H{"Data": assetDetails1})

// 	} else {
// 		res, err := http.Get("http://localhost:8080/Users")
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		defer res.Body.Close()

// 		var responseData map[string][]map[string]interface{}
// 		err = json.NewDecoder(res.Body).Decode(&responseData)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}

// 		// Extract asset details from the response
// 		assetDetails, ok := responseData["data"]
// 		if !ok {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data format error"})
// 			return
// 		}

// 		// Pass assetDetails to the HTML template
// 		c.HTML(http.StatusOK, "admin.html", gin.H{"Data": assetDetails})
// 	}
// }

func GetUsers(c *gin.Context) {
	if result.Type == "User" {
		employeeID := result.EmpID
		url := fmt.Sprintf("http://localhost:8080/admin/%s", employeeID)
		res1, err := http.Get(url)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer res1.Body.Close()

		var responseData1 map[string]interface{}
		err = json.NewDecoder(res1.Body).Decode(&responseData1)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		assetDetails1, ok := responseData1["data"].(map[string]interface{})
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data format error"})
			return
		}
		

		// Create an instance of AssetDetail
		assetDetail := model.AssetDetail{}
		if empID, ok := assetDetails1["empID"].(string); ok {
			assetDetail.EmpID = empID
		}

		// Extract and convert image URLs
		if laptop, ok := assetDetails1["laptop"].(map[string]interface{}); ok {
			if imageurl, ok := laptop["imageurl"].(string); ok {
				assetDetail.Laptop.ImageBase64 = []byte(imageurl)
			}
			if imageName, ok := laptop["imageName"].(string); ok {
				assetDetail.Laptop.ImageName = imageName
			}
		}

		if mouse, ok := assetDetails1["mouse"].(map[string]interface{}); ok {
			if imageurl, ok := mouse["imageurl"].(string); ok {
				assetDetail.Mouse.ImageBase64 = []byte(imageurl)
			}
			if imageName, ok := mouse["imageName"].(string); ok {
				assetDetail.Mouse.ImageName = imageName
			}
		}

		if headphones, ok := assetDetails1["headphones"].(map[string]interface{}); ok {
			if imageurl, ok := headphones["imageurl"].(string); ok {
				assetDetail.Headphones.ImageBase64 = []byte(imageurl)
			}
			if imageName, ok := headphones["imageName"].(string); ok {
				assetDetail.Headphones.ImageName = imageName
			}
		}

		// Convert image data to base64
		laptopImageBase64 := base64.StdEncoding.EncodeToString(assetDetail.Laptop.ImageBase64)
		mouseImageBase64 := base64.StdEncoding.EncodeToString(assetDetail.Mouse.ImageBase64)
		headphonesImageBase64 := base64.StdEncoding.EncodeToString(assetDetail.Headphones.ImageBase64)

		// Pass asset details and images to the HTML template
		c.HTML(http.StatusOK, "user.html", gin.H{
			"Data":                  assetDetails1,
			"LaptopImageBase64":     laptopImageBase64,
			"LaptopImageName":       assetDetail.Laptop.ImageName,
			"MouseImageBase64":      mouseImageBase64,
			"MouseImageName":        assetDetail.Mouse.ImageName,
			"HeadphonesImageBase64": headphonesImageBase64,
			"HeadphonesImageName":   assetDetail.Headphones.ImageName,
		})

	} else {
		res, err := http.Get("http://localhost:8080/Users")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		defer res.Body.Close()

		var responseData map[string][]map[string]interface{}
		err = json.NewDecoder(res.Body).Decode(&responseData)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		assetDetails, ok := responseData["data"]
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Data format error"})
			return
		}

		// Pass asset details to the HTML template
		c.HTML(http.StatusOK, "admin.html", gin.H{"Data": assetDetails})
	}
}
