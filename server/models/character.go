package models

import "time"

// Character represents a character in the game
type Character struct {
	ID          int       `json:"id" db:"id"`
	UserID      int       `json:"user_id" db:"user_id"`
	Name        string    `json:"name" db:"name"`
	Level       int       `json:"level" db:"level"`
	Experience  int       `json:"experience" db:"experience"`
	Health      int       `json:"health" db:"health"`
	MaxHealth   int       `json:"max_health" db:"max_health"`
	Mana        int       `json:"mana" db:"mana"`
	MaxMana     int       `json:"max_mana" db:"max_mana"`
	Attack      int       `json:"attack" db:"attack"`
	Defense     int       `json:"defense" db:"defense"`
	Speed       int       `json:"speed" db:"speed"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CharacterCreation represents data for creating a new character
type CharacterCreation struct {
	Name string `json:"name" binding:"required,min=1,max=30"`
}

// CharacterStats represents the stats of a character
type CharacterStats struct {
	Level      int `json:"level"`
	Experience int `json:"experience"`
	Health     int `json:"health"`
	MaxHealth  int `json:"max_health"`
	Mana       int `json:"mana"`
	MaxMana    int `json:"max_mana"`
	Attack     int `json:"attack"`
	Defense    int `json:"defense"`
	Speed      int `json:"speed"`
}
