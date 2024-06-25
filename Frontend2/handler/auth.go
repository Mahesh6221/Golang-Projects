package handler

import (
	"FRONTEND2/model"
	"bytes"
	"encoding/json"
	// "fmt"

	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)
}

var result model.User

func PostLogin(c *gin.Context) {
	mailID := c.PostForm("mailID")
	password := c.PostForm("password")

	var user model.User

	user.MailID = mailID
	user.Password = password

	jsonData, err := json.Marshal(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", "http://localhost:8080/login", bytes.NewBuffer(jsonData))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error "})
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error sending request"})
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		var errorResponse struct {
			Error string `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding response"})
			return
		}
		errorMessage := errorResponse.Error
		c.HTML(resp.StatusCode, "login.html", gin.H{"error": errorMessage})
		return
	}

	if resp.StatusCode != http.StatusOK {
		c.HTML(resp.StatusCode, "login.html", gin.H{"error": "Login failed"})
		return
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error decoding response"})
		return
	}
	// fmt.Printf("result.Role: %v\n", result.Role)
	c.SetCookie("token", result.Token, 1800, "/", "http://localhost:9090", false, true) // Expires in 2 hours
	c.SetCookie("refresh_token", result.Refresh_token, 1800, "/", "http://localhost:9090", false, true)
	// fmt.Printf("result: %v\n", result)
	if result.MailID != "" {
		c.Redirect(301, "/users")
		return
	}
	c.HTML(http.StatusOK, "login.html", gin.H{})
}
func CheckTokenExpiration() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := c.Cookie("token")
		if err != nil {
			c.Redirect(http.StatusFound, "/login")
			return
		}
		c.Next()
	}
}
func UserLogout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.Redirect(301, "/login")
}
