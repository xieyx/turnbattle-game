#!/bin/bash

# Test authentication APIs
echo "Testing TurnBattle Authentication APIs"
echo "====================================="

# Start the server in the background
echo "Starting server..."
cd server
go run main.go > server.log 2>&1 &
SERVER_PID=$!
cd ..

# Wait for server to start
sleep 3

# Test registration
echo "1. Testing user registration..."
curl -X POST http://localhost:8080/api/v1/users/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }' | jq .

echo -e "\n"

# Test login
echo "2. Testing user login..."
curl -X POST http://localhost:8080/api/v1/users/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }' | jq . > token.json

# Extract token
TOKEN=$(jq -r '.token' token.json)
echo "Token: $TOKEN"

echo -e "\n"

# Test profile with token
echo "3. Testing user profile with token..."
curl -X GET http://localhost:8080/api/v1/users/profile \
  -H "Authorization: Bearer $TOKEN" | jq .

echo -e "\n"

# Test profile without token
echo "4. Testing user profile without token..."
curl -X GET http://localhost:8080/api/v1/users/profile | jq .

echo -e "\n"

# Stop the server
echo "Stopping server..."
kill $SERVER_PID

echo "Test completed!"
