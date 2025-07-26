#!/bin/bash

# Test battle APIs
echo "Testing TurnBattle Battle APIs"
echo "==============================="

# Start the server in the background
echo "Starting server..."
cd server
go run main.go > server.log 2>&1 &
SERVER_PID=$!
cd ..

# Wait for server to start
sleep 3

# Test battle creation
echo "1. Testing battle creation..."
curl -X POST http://localhost:8080/api/v1/battles \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy_token" \
  -d '{
    "player2_id": 2
  }' | jq .

echo -e "\n"

# Test battle start
echo "2. Testing battle start..."
curl -X POST http://localhost:8080/api/v1/battles/1/start \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy_token" | jq .

echo -e "\n"

# Test battle turn execution
echo "3. Testing battle turn execution..."
curl -X POST http://localhost:8080/api/v1/battles/1/turns \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer dummy_token" \
  -d '{
    "actor_id": 1,
    "action": "attack",
    "target_id": 2
  }' | jq .

echo -e "\n"

# Test battle retrieval
echo "4. Testing battle retrieval..."
curl -X GET http://localhost:8080/api/v1/battles/1 \
  -H "Authorization: Bearer dummy_token" | jq .

echo -e "\n"

# Stop the server
echo "Stopping server..."
kill $SERVER_PID

echo "Battle API test completed!"
