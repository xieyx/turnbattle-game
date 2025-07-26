package models

import "time"

// Skill represents a skill in the game
type Skill struct {
	ID          int       `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Type        string    `json:"type" db:"type"` // "attack", "heal", "buff", "debuff"
	Damage      int       `json:"damage" db:"damage"`
	ManaCost    int       `json:"mana_cost" db:"mana_cost"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// CharacterSkill represents a skill that a character has learned
type CharacterSkill struct {
	ID          int       `json:"id" db:"id"`
	CharacterID int       `json:"character_id" db:"character_id"`
	SkillID     int       `json:"skill_id" db:"skill_id"`
	Level       int       `json:"level" db:"level"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// SkillEffect represents the effect of a skill
type SkillEffect struct {
	ID          int       `json:"id" db:"id"`
	SkillID     int       `json:"skill_id" db:"skill_id"`
	EffectType  string    `json:"effect_type" db:"effect_type"` // "damage", "heal", "buff", "debuff"
	Value       int       `json:"value" db:"value"`
	Duration    int       `json:"duration" db:"duration"` // 持续回合数
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
