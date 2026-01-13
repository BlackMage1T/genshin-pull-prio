package main

import (
	"os"

	"github.com/BlackMage1T/genshin-pull-prio/internal/database"
	"github.com/BlackMage1T/genshin-pull-prio/internal/handlers"
	"github.com/BlackMage1T/genshin-pull-prio/internal/models"

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
	r.GET("/debug/stats", func(c *gin.Context) {
		var charCount, talentCount int64
		db.Model(&models.Character{}).Count(&charCount)
		db.Model(&models.Talent{}).Count(&talentCount)

		c.JSON(200, gin.H{
			"characters_in_db": charCount,
			"talents_in_db":    talentCount,
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default for local testing
	}
	r.Run(":" + port)
}
