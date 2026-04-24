package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Beast of the Hidden Leaf", Artist: "Juice World", Price: 55.80},
	{ID: "2", Title: "The Dark Knight", Artist: "Batman", Price: 90.45},
	{ID: "3", Title: "The Kryptonian", Artist: "Superman", Price: 80.35},
	{ID: "4", Title: "The Honoured One", Artist: "Gojo Satoru", Price: 500.80},
}

func main() {
	router := gin.Default()

	// Routes
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumByID) // NEW: Search for one album
	router.GET("/profile", getimages)

	router.Run("localhost:8080")
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range albums {
		if a.ID == id {
			c.JSON(http.StatusOK, a)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func getimages(c *gin.Context) {
	email := c.Query("email")
	email = strings.ToLower(strings.TrimSpace(email))

	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "email query parameter is required",
		})
		return
	}
	fmt.Printf("[LOG] %s: New request for avatar: %s\n", time.Now().Format("2006-01-02"), email)
	hash := md5.Sum([]byte(email))
	hashstr := hex.EncodeToString(hash[:])

	gravataURL := fmt.Sprintf("https://www.gravatar.com/avatar/%s?s=200", hashstr)
	dicebearURL := fmt.Sprintf("https://api.dicebear.com/9.x/bottts/svg?seed=%s", email)

	c.JSON(http.StatusOK, gin.H{
		"email":        email,
		"human_avatar": gravataURL,
		"robot_avatar": dicebearURL,
		"developer":    "Finix",
		"github":       "https://github.com/angelgabriel-dotcom/go-avatar-api",
	})
}
