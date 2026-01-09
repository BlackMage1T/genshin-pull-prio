package models

import "gorm.io/gorm"

type Character struct {
	gorm.Model
	Name           string   `json:"name" gorm:"uniqueIndex"`
	Title          string   `json:"title"`
	Vision         string   `json:"vision"`
	Weapon         string   `json:"weapon"`
	Gender         string   `json:"gender"`
	Nation         string   `json:"nation"`
	Affiliation    string   `json:"affiliation"`
	Rarity         int      `json:"rarity"`
	SkillTalents   []Talent `json:"skillTalents" gorm:"foreignKey:CharacterID"`
	PassiveTalents []Talent `json:"passiveTalents" gorm:"foreignKey:CharacterID"`
	Constellations []Talent `json:"constellations" gorm:"foreignKey:CharacterID"`
}

type Talent struct {
	gorm.Model
	CharacterID uint
	Name        string         `json:"name"`
	Unlock      string         `json:"unlock"`
	Description string         `json:"description"`
	Upgrades    []UpgradeLevel `json:"upgrades" gorm:"foreignKey:TalentID"`
	SkillType   string         `json:"type"`
	Level       int            `json:"level"`
}

type UpgradeLevel struct {
	gorm.Model
	TalentID uint
	Name     string `json:"name"`
	Value    string `json:"value"`
}
