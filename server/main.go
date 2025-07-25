package main

import (
"log"
"net/http"

"github.com/gin-gonic/gin"
"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 允许所有来源，生产环境中应该更严格
	},
}

func main() {
	// 设置Gin路由
	router := gin.Default()

	// 基础路由
	router.GET("/", func(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"message": "Welcome to TurnBattle API",
"version": "1.0.0",
})
})

	// WebSocket路由
	router.GET("/ws", func(c *gin.Context) {
// 升级HTTP连接到WebSocket
conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
if err != nil {
log.Printf("WebSocket upgrade failed: %v", err)
c.JSON(http.StatusInternalServerError, gin.H{"error": "WebSocket upgrade failed"})
return
}
defer conn.Close()

		// 处理WebSocket连接
		for {
			// 读取消息
			messageType, message, err := conn.ReadMessage()
			if err != nil {
				log.Printf("WebSocket read error: %v", err)
				break
			}

			// 回显消息
			log.Printf("Received: %s", message)
			err = conn.WriteMessage(messageType, message)
			if err != nil {
				log.Printf("WebSocket write error: %v", err)
				break
			}
		}
	})

	// 用户相关路由
	router.POST("/api/v1/users/register", func(c *gin.Context) {
c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
})

	router.POST("/api/v1/users/login", func(c *gin.Context) {
c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
})

	// 战斗相关路由
	router.POST("/api/v1/battles", func(c *gin.Context) {
c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
})

	router.GET("/api/v1/battles/:id", func(c *gin.Context) {
c.JSON(http.StatusNotImplemented, gin.H{"error": "Not implemented yet"})
})

	// 启动服务器
	log.Println("Starting TurnBattle server on :8080")
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
