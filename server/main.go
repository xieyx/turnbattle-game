package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	_ "github.com/lib/pq"

	"github.com/xieyx/turnbattle-game/server/middleware"
	"github.com/xieyx/turnbattle-game/server/models"
	"github.com/xieyx/turnbattle-game/server/services"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源，生产环境中应该更严格
	},
}

// Global database connection
var db *sql.DB

func main() {
	// Initialize database connection
	// 注意：在生产环境中，应该从环境变量或配置文件中读取数据库连接信息
	var err error
	db, err = sql.Open("postgres", "user=turnbattle_user password=password dbname=turnbattle sslmode=disable")
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	// Test the database connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	// Initialize services
	userService := services.NewUserService(db)
	battleService := services.NewBattleService(db)

	// Set up Gin router
	router := gin.Default()

	// Public routes
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to TurnBattle API",
			"version": "1.0.0",
		})
	})

	// User authentication routes
	router.POST("/api/v1/users/register", func(c *gin.Context) {
		var registration models.UserRegistration
		if err := c.ShouldBindJSON(&registration); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		user, err := userService.RegisterUser(&registration)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, user)
	})

	router.POST("/api/v1/users/login", func(c *gin.Context) {
		var login models.UserLogin
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		token, err := userService.AuthenticateUser(&login)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"token": token})
	})

	// Protected routes (require authentication)
	authorized := router.Group("/")
	authorized.Use(middleware.JWTAuthMiddleware())
	{
		authorized.GET("/api/v1/users/profile", func(c *gin.Context) {
			userID, exists := c.Get("user_id")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
				return
			}

			profile, err := userService.GetUserProfile(userID.(int))
			if err != nil {
				c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, profile)
		})

		// Character routes
		authorized.POST("/api/v1/characters", func(c *gin.Context) {
			userID, exists := c.Get("user_id")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
				return
			}

			var creation models.CharacterCreation
			if err := c.ShouldBindJSON(&creation); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// TODO: Implement character creation
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
		})

		authorized.GET("/api/v1/characters", func(c *gin.Context) {
			userID, exists := c.Get("user_id")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
				return
			}

			// TODO: Implement character listing
			c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
		})

		// Battle routes
		authorized.POST("/api/v1/battles", func(c *gin.Context) {
			userID, exists := c.Get("user_id")
			if !exists {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "User ID not found in context"})
				return
			}

			var battleRequest struct {
				Player2ID int `json:"player2_id" binding:"required"`
			}
			if err := c.ShouldBindJSON(&battleRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// TODO: Get player1 character ID from user ID
			player1ID := 1 // Placeholder

			battle, err := battleService.CreateBattle(player1ID, battleRequest.Player2ID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusCreated, battle)
		})

		authorized.POST("/api/v1/battles/:id/start", func(c *gin.Context) {
			battleID, err := c.Params.Get("id")
			if !err {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Battle ID is required"})
				return
			}

			// Convert battleID to int
			// TODO: Implement proper conversion
			id := 1 // Placeholder

			err = battleService.StartBattle(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, gin.H{"message": "Battle started"})
		})

		authorized.POST("/api/v1/battles/:id/turns", func(c *gin.Context) {
			battleID, err := c.Params.Get("id")
			if !err {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Battle ID is required"})
				return
			}

			var turnRequest struct {
				ActorID  int    `json:"actor_id" binding:"required"`
				Action   string `json:"action" binding:"required"`
				TargetID *int   `json:"target_id"`
			}
			if err := c.ShouldBindJSON(&turnRequest); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}

			// Convert battleID to int
			// TODO: Implement proper conversion
			id := 1 // Placeholder

			turn, err := battleService.ExecuteTurn(id, turnRequest.ActorID, turnRequest.Action, turnRequest.TargetID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, turn)
		})

		authorized.GET("/api/v1/battles/:id", func(c *gin.Context) {
			battleID, err := c.Params.Get("id")
			if !err {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Battle ID is required"})
				return
			}

			// Convert battleID to int
			// TODO: Implement proper conversion
			id := 1 // Placeholder

			battle, err := battleService.GetBattle(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			participants, err := battleService.GetBattleParticipants(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			turns, err := battleService.GetBattleTurns(id)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			response := struct {
				Battle      *models.Battle              `json:"battle"`
				Participants []models.BattleParticipant `json:"participants"`
				Turns       []models.BattleTurn         `json:"turns"`
			}{
				Battle:      battle,
				Participants: participants,
				Turns:       turns,
			}

			c.JSON(http.StatusOK, response)
		})
	}

	// WebSocket route
	router.GET("/ws", func(c *gin.Context) {
		// Upgrade HTTP connection to WebSocket
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Printf("WebSocket upgrade failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
			return
		}
		defer conn.Close()

		// Handle WebSocket connection
		for {
			// Read message
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("WebSocket read error: %v", err)
				break
			}

			// Echo message
			log.Printf("Received: %s", message)
			err = conn.WriteMessage(messageType, message)
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				break
			}
		}
	})

	// Start server
	log.Println("Starting TurnBattle server on :8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}