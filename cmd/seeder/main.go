package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/BlackMage1T/genshin-pull-prio/internal/database"
	"github.com/BlackMage1T/genshin-pull-prio/internal/models"
)

func main() {
	db := database.InitDB()

	// Path to your JSON files
	seedDir := "internal/database/seed_data/char"

	// 1. Read the directory
	files, err := os.ReadDir(seedDir)
	if err != nil {
		log.Fatalf("Could not read seed directory: %v", err)
	}

	fmt.Println("--- Starting Seeder ---")

	for _, file := range files {
		if filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(seedDir, file.Name())

			// 2. Read the file
			content, err := os.ReadFile(filePath)
			if err != nil {
				log.Printf("Error reading %s: %v", file.Name(), err)
				continue
			}

			// 3. Parse JSON into struct
			var character models.Character
			if err := json.Unmarshal(content, &character); err != nil {
				log.Printf("Error parsing %s: %v", file.Name(), err)
				continue
			}

			// 4. Save to Database
			// GORM automatically saves nested Talents if the struct is populated
			if err := db.Create(&character).Error; err != nil {
				log.Printf("Error saving %s to DB: %v", character.Name, err)
				continue
			}

			fmt.Printf("Successfully seeded: %s\n", character.Name)
		}
	}

	fmt.Println("--- Seeding Complete ---")
}
