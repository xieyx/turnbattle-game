package services

import (
"database/sql"
"errors"
"log"
"time"

"github.com/xieyx/turnbattle-game/server/models"
"github.com/xieyx/turnbattle-game/server/utils"
)

// UserService handles user-related business logic
type UserService struct {
	db *sql.DB
}

// NewUserService creates a new UserService
func NewUserService(db *sql.DB) *UserService {
	return &UserService{db: db}
}

// RegisterUser registers a new user
func (s *UserService) RegisterUser(registration *models.UserRegistration) (*models.User, error) {
	// Check if username or email already exists
	var exists bool
	err := s.db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE username = $1 OR email = $2)", 
registration.Username, registration.Email).Scan(&exists)
	if err != nil {
		return nil, err
	}
	
	if exists {
		return nil, errors.New("username or email already exists")
	}
	
	// Hash the password
	hashedPassword, err := utils.HashPassword(registration.Password)
	if err != nil {
		return nil, err
	}
	
	// Create the user
	user := &models.User{
		Username:  registration.Username,
		Email:     registration.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	err = s.db.QueryRow(
"INSERT INTO users (username, email, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id",
user.Username, user.Email, user.Password, user.CreatedAt, user.UpdatedAt,
).Scan(&user.ID)
	
	if err != nil {
		return nil, err
	}
	
	// Don't return the password
user.Password = ""

return user, nil
}

// AuthenticateUser authenticates a user and returns a JWT token
func (s *UserService) AuthenticateUser(login *models.UserLogin) (string, error) {
// Get user from database
var user models.User
err := s.db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", 
login.Username).Scan(&user.ID, &user.Username, &user.Password)
if err != nil {
if err == sql.ErrNoRows {
return "", errors.New("invalid username or password")
}
return "", err
}

// Check password
err = utils.CheckPassword(login.Password, user.Password)
if err != nil {
return "", errors.New("invalid username or password")
}

// Generate JWT token
token, err := utils.GenerateToken(user.ID, user.Username)
if err != nil {
log.Printf("Failed to generate token: %v", err)
return "", errors.New("failed to generate token")
}

return token, nil
}

// GetUserProfile gets a user's profile by ID
func (s *UserService) GetUserProfile(userID int) (*models.UserProfile, error) {
	var profile models.UserProfile
	err := s.db.QueryRow("SELECT id, username, email, created_at FROM users WHERE id = $1", 
userID).Scan(&profile.ID, &profile.Username, &profile.Email, &profile.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &profile, nil
}
