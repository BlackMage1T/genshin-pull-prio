package main

import (
	"github.com/BlackMage1T/genshin-pull-prio/internal/database"
	"github.com/BlackMage1T/genshin-pull-prio/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Initialize DB
	db := database.InitDB()

	r := gin.Default()

	// 2. Pass the 'db' variable into the handler function
	// This resolves the "undefined" error!
	h := handlers.GetCharacters(db)

	r.GET("/characters", h)       // Returns list
	r.GET("/characters/:name", h) // Returns specific char

	r.Run(":8080")
}
