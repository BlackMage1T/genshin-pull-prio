package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/BlackMage1T/genshin-pull-prio/internal/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetCharacters now accepts the db and RETURNS a function Gin can use
func GetCharacters(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		name := strings.Title(strings.ToLower(c.Param("name")))

		if name == "" {
			var characters []models.Character
			err := db.Preload("SkillTalents.Upgrades").
				Preload("PassiveTalents").
				Preload("Constellations").
				Find(&characters).Error
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, characters)
			return
		}

		var character models.Character
		err := db.Preload("SkillTalents.Upgrades").
			Preload("PassiveTalents").
			Preload("Constellations").
			Where("name = ?", name).
			First(&character).Error

		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Character not found"})
			return
		}

		fmt.Printf("Database Record Found: %+v\n", character)

		c.JSON(http.StatusOK, character)
	}
}
