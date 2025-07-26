package models

import "time"

// Battle represents a battle between characters
type Battle struct {
	ID            int       `json:"id" db:"id"`
	Player1ID     int       `json:"player1_id" db:"player1_id"`
	Player2ID     int       `json:"player2_id" db:"player2_id"`
	WinnerID      *int      `json:"winner_id" db:"winner_id"`
	Status        string    `json:"status" db:"status"` // "pending", "in_progress", "completed"
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// BattleTurn represents a turn in a battle
type BattleTurn struct {
	ID        int       `json:"id" db:"id"`
	BattleID  int       `json:"battle_id" db:"battle_id"`
	TurnNumber int      `json:"turn_number" db:"turn_number"`
	ActorID   int       `json:"actor_id" db:"actor_id"`
	Action    string    `json:"action" db:"action"` // "attack", "skill", "defend", "item"
	TargetID  *int      `json:"target_id" db:"target_id"`
	Result    string    `json:"result" db:"result"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// BattleParticipant represents a participant in a battle
type BattleParticipant struct {
	ID          int       `json:"id" db:"id"`
	BattleID    int       `json:"battle_id" db:"battle_id"`
	CharacterID int       `json:"character_id" db:"character_id"`
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
