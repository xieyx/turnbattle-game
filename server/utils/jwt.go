package utils

import (
"errors"
"time"

"github.com/golang-jwt/jwt/v4"
)

var (
// JWTSecret is the secret key for JWT signing
JWTSecret = []byte("turnbattle_secret_key_2025")

// ErrInvalidToken is returned when the token is invalid
ErrInvalidToken = errors.New("invalid token")

// ErrExpiredToken is returned when the token is expired
ErrExpiredToken = errors.New("expired token")
)

// JWTClaims represents the claims in a JWT
type JWTClaims struct {
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateToken generates a new JWT token for a user
func GenerateToken(userID int, username string) (string, error) {
	claims := &JWTClaims{
		UserID:   userID,
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 24小时过期
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "turnbattle",
		},
	}
	
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(JWTSecret)
}

// VerifyToken verifies a JWT token and returns the claims
func VerifyToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return JWTSecret, nil
	})
	
	if err != nil {
		return nil, ErrInvalidToken
	}
	
	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}
	
	// 检查是否过期
	if claims.ExpiresAt != nil && claims.ExpiresAt.Before(time.Now()) {
		return nil, ErrExpiredToken
	}
	
	return claims, nil
}
