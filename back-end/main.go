package main

import (
	"fmt"
	"log"
	_ "path/filepath"
	"tic-tac-toe/internal/config"
	"tic-tac-toe/internal/handlers"
	"tic-tac-toe/internal/websocket"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()
	hub := websocket.NewHub()
	go hub.Run()

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := handlers.NewHandler()
	router := handler.SetupRoutes()

	router.GET("/ws", func(c *gin.Context) {
		websocket.ServeWs(hub, c.Writer, c.Request)
	})

	fmt.Printf("Server starting on %s:%s\n", cfg.Host, cfg.Port)
	if cfg.Debug {
		fmt.Println("Debug mode enabled")
	}

	log.Fatal(router.Run(":" + cfg.Port))
}
