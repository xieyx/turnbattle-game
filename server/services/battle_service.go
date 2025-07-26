package services

import (
	"database/sql"
	"errors"
	"log"
	"math/rand"
	"time"

	"github.com/xieyx/turnbattle-game/server/models"
)

// BattleService handles battle-related business logic
type BattleService struct {
	db *sql.DB
}

// NewBattleService creates a new BattleService
func NewBattleService(db *sql.DB) *BattleService {
	return &BattleService{db: db}
}

// CreateBattle creates a new battle between two characters
func (s *BattleService) CreateBattle(player1ID, player2ID int) (*models.Battle, error) {
	// Check if both characters exist
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM characters WHERE id IN ($1, $2))",
		player1ID, player2ID).Scan(&exists)
	if err != nil {
		return nil, err
	}

	if !exists {
		return nil, errors.New("one or both characters do not exist")
	}

	// Create the battle
	battle := &models.Battle{
		Player1ID: player1ID,
		Player2ID: player2ID,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	err = s.db.QueryRow(
		"INSERT INTO battles (player1_id, player2_id, status, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		battle.Player1ID, battle.Player2ID, battle.Status, battle.CreatedAt, battle.UpdatedAt,
	).Scan(&battle.ID)

	if err != nil {
		return nil, err
	}

	return battle, nil
}

// StartBattle starts a battle and initializes participants
func (s *BattleService) StartBattle(battleID int) error {
	// Get battle
	var battle models.Battle
	err := s.db.QueryRow("SELECT id, player1_id, player2_id, status FROM battles WHERE id = $1",
		battleID).Scan(&battle.ID, &battle.Player1ID, &battle.Player2ID, &battle.Status)
	if err != nil {
		return err
	}

	if battle.Status != "pending" {
		return errors.New("battle is not in pending status")
	}

	// Get character stats for both players
	char1, err := s.getCharacterStats(battle.Player1ID)
	if err != nil {
		return err
	}

	char2, err := s.getCharacterStats(battle.Player2ID)
	if err != nil {
		return err
	}

	// Create battle participants
	_, err = s.db.Exec(
		"INSERT INTO battle_participants (battle_id, character_id, health, max_health, mana, max_mana, attack, defense, speed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		battleID, battle.Player1ID, char1.Health, char1.MaxHealth, char1.Mana, char1.MaxMana, char1.Attack, char1.Defense, char1.Speed, time.Now(), time.Now(),
	)
	if err != nil {
		return err
	}

	_, err = s.db.Exec(
		"INSERT INTO battle_participants (battle_id, character_id, health, max_health, mana, max_mana, attack, defense, speed, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)",
		battleID, battle.Player2ID, char2.Health, char2.MaxHealth, char2.Mana, char2.MaxMana, char2.Attack, char2.Defense, char2.Speed, time.Now(), time.Now(),
	)
	if err != nil {
		return err
	}

	// Update battle status
	_, err = s.db.Exec("UPDATE battles SET status = 'in_progress', updated_at = $1 WHERE id = $2",
		time.Now(), battleID)

	return err
}

// ExecuteTurn executes a turn in a battle
func (s *BattleService) ExecuteTurn(battleID, actorID int, action string, targetID *int) (*models.BattleTurn, error) {
	// Get battle
	var battle models.Battle
	err := s.db.QueryRow("SELECT id, status FROM battles WHERE id = $1",
		battleID).Scan(&battle.ID, &battle.Status)
	if err != nil {
		return nil, err
	}

	if battle.Status != "in_progress" {
		return nil, errors.New("battle is not in progress")
	}

	// Get the next turn number
	var turnNumber int
	err = s.db.QueryRow("SELECT COALESCE(MAX(turn_number), 0) + 1 FROM battle_turns WHERE battle_id = $1",
		battleID).Scan(&turnNumber)
	if err != nil {
		return nil, err
	}

	// Create battle turn
	battleTurn := &models.BattleTurn{
		BattleID:   battleID,
		TurnNumber: turnNumber,
		ActorID:    actorID,
		Action:     action,
		TargetID:   targetID,
		CreatedAt:  time.Now(),
	}

	// Execute the action and set the result
	result, err := s.executeAction(battleID, actorID, action, targetID)
	if err != nil {
		return nil, err
	}

	battleTurn.Result = result

	// Save the turn
	err = s.db.QueryRow(
		"INSERT INTO battle_turns (battle_id, turn_number, actor_id, action, target_id, result, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		battleTurn.BattleID, battleTurn.TurnNumber, battleTurn.ActorID, battleTurn.Action, battleTurn.TargetID, battleTurn.Result, battleTurn.CreatedAt,
	).Scan(&battleTurn.ID)

	if err != nil {
		return nil, err
	}

	// Check if the battle is over
	winnerID, err := s.checkBattleEnd(battleID)
	if err != nil {
		return nil, err
	}

	if winnerID != nil {
		// Battle is over, update battle status
		_, err = s.db.Exec("UPDATE battles SET status = 'completed', winner_id = $1, updated_at = $2 WHERE id = $3",
			winnerID, time.Now(), battleID)
		if err != nil {
			return nil, err
		}
	}

	return battleTurn, nil
}

// getCharacterStats gets the stats of a character
func (s *BattleService) getCharacterStats(characterID int) (*models.CharacterStats, error) {
	var stats models.CharacterStats
	err := s.db.QueryRow(
		"SELECT level, experience, health, max_health, mana, max_mana, attack, defense, speed FROM characters WHERE id = $1",
		characterID,
	).Scan(&stats.Level, &stats.Experience, &stats.Health, &stats.MaxHealth, &stats.Mana, &stats.MaxMana, &stats.Attack, &stats.Defense, &stats.Speed)

	if err != nil {
		return nil, err
	}

	return &stats, nil
}

// executeAction executes an action and returns the result
func (s *BattleService) executeAction(battleID, actorID int, action string, targetID *int) (string, error) {
	switch action {
	case "attack":
		return s.executeAttack(battleID, actorID, *targetID)
	case "defend":
		return s.executeDefend(battleID, actorID)
	default:
		return "Invalid action", nil
	}
}

// executeAttack executes an attack action
func (s *BattleService) executeAttack(battleID, actorID, targetID int) (string, error) {
	// Get actor stats
	var actorAttack, actorSpeed int
	err := s.db.QueryRow(
		"SELECT attack, speed FROM battle_participants WHERE battle_id = $1 AND character_id = $2",
		battleID, actorID,
	).Scan(&actorAttack, &actorSpeed)
	if err != nil {
		return "", err
	}

	// Get target stats
	var targetHealth, targetDefense int
	err = s.db.QueryRow(
		"SELECT health, defense FROM battle_participants WHERE battle_id = $1 AND character_id = $2",
		battleID, targetID,
	).Scan(&targetHealth, &targetDefense)
	if err != nil {
		return "", err
	}

	// Calculate damage
	damage := actorAttack - targetDefense/2
	if damage < 1 {
		damage = 1
	}

	// Apply critical hit chance
	if rand.Intn(100) < 10 { // 10% chance
		damage *= 2
	}

	// Update target health
	newHealth := targetHealth - damage
	if newHealth < 0 {
		newHealth = 0
	}

	_, err = s.db.Exec(
		"UPDATE battle_participants SET health = $1, updated_at = $2 WHERE battle_id = $3 AND character_id = $4",
		newHealth, time.Now(), battleID, targetID,
	)
	if err != nil {
		return "", err
	}

	result := ""
	if newHealth <= 0 {
		result = "Attack successful! Target defeated!"
	} else {
		result = "Attack successful! Dealt " + string(rune(damage)) + " damage."
	}

	return result, nil
}

// executeDefend executes a defend action
func (s *BattleService) executeDefend(battleID, actorID int) (string, error) {
	// Increase defense temporarily (in a real implementation, this would be more complex)
	_, err := s.db.Exec(
		"UPDATE battle_participants SET defense = defense * 1.5, updated_at = $1 WHERE battle_id = $2 AND character_id = $3",
		time.Now(), battleID, actorID,
	)
	if err != nil {
		return "", err
	}

	return "Defense increased!", nil
}

// checkBattleEnd checks if the battle is over and returns the winner ID
func (s *BattleService) checkBattleEnd(battleID int) (*int, error) {
	// Get participants
	rows, err := s.db.Query(
		"SELECT character_id, health FROM battle_participants WHERE battle_id = $1",
		battleID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []struct {
		CharacterID int
		Health      int
	}

	for rows.Next() {
		var p struct {
			CharacterID int
			Health      int
		}
		err := rows.Scan(&p.CharacterID, &p.Health)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	// Check if any participant has 0 health
	for _, p := range participants {
		if p.Health <= 0 {
			// Find the other participant as the winner
			for _, p2 := range participants {
				if p2.CharacterID != p.CharacterID {
					return &p2.CharacterID, nil
				}
			}
		}
	}

	return nil, nil
}

// GetBattle gets a battle by ID
func (s *BattleService) GetBattle(battleID int) (*models.Battle, error) {
	var battle models.Battle
	err := s.db.QueryRow(
		"SELECT id, player1_id, player2_id, winner_id, status, created_at, updated_at FROM battles WHERE id = $1",
		battleID,
	).Scan(&battle.ID, &battle.Player1ID, &battle.Player2ID, &battle.WinnerID, &battle.Status, &battle.CreatedAt, &battle.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &battle, nil
}

// GetBattleParticipants gets the participants of a battle
func (s *BattleService) GetBattleParticipants(battleID int) ([]models.BattleParticipant, error) {
	rows, err := s.db.Query(
		"SELECT id, battle_id, character_id, health, max_health, mana, max_mana, attack, defense, speed, created_at, updated_at FROM battle_participants WHERE battle_id = $1",
		battleID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var participants []models.BattleParticipant
	for rows.Next() {
		var p models.BattleParticipant
		err := rows.Scan(&p.ID, &p.BattleID, &p.CharacterID, &p.Health, &p.MaxHealth, &p.Mana, &p.MaxMana, &p.Attack, &p.Defense, &p.Speed, &p.CreatedAt, &p.UpdatedAt)
		if err != nil {
			return nil, err
		}
		participants = append(participants, p)
	}

	return participants, nil
}

// GetBattleTurns gets the turns of a battle
func (s *BattleService) GetBattleTurns(battleID int) ([]models.BattleTurn, error) {
	rows, err := s.db.Query(
		"SELECT id, battle_id, turn_number, actor_id, action, target_id, result, created_at FROM battle_turns WHERE battle_id = $1 ORDER BY turn_number",
		battleID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var turns []models.BattleTurn
	for rows.Next() {
		var t models.BattleTurn
		err := rows.Scan(&t.ID, &t.BattleID, &t.TurnNumber, &t.ActorID, &t.Action, &t.TargetID, &t.Result, &t.CreatedAt)
		if err != nil {
			return nil, err
		}
		turns = append(turns, t)
	}

	return turns, nil
}