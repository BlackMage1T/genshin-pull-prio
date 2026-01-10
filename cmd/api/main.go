package main

import (
	"os"

	"github.com/BlackMage1T/genshin-pull-prio/internal/database"
	"github.com/BlackMage1T/genshin-pull-prio/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize DB
	db := database.InitDB()

	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 2. Pass the 'db' variable into the handler function
	// This resolves the "undefined" error!
	h := handlers.GetCharacters(db)

	r.GET("/characters", h)       // Returns list
	r.GET("/characters/:name", h) // Returns specific char

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local testing
	}
	r.Run(":" + port)
}
