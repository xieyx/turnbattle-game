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

	// Battle routes (not implemented yet)
	router.POST("/api/v1/battles", func(c *gin.Context) {
c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
})

	router.GET("/api/v1/battles/:id", func(c *gin.Context) {
c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
})

	// Start server
	log.Println("Starting TurnBattle server on :8080")
	err = router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
