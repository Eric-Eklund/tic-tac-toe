package main

import (
	"fmt"
	"log"
	_ "path/filepath"
	"tic-tac-toe/internal/config"
	"tic-tac-toe/internal/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	if !cfg.Debug {
		gin.SetMode(gin.ReleaseMode)
	}

	handler := handlers.NewHandler()
	router := handler.SetupRoutes()

	fmt.Printf("Server starting on %s:%s\n", cfg.Host, cfg.Port)
	if cfg.Debug {
		fmt.Println("Debug mode enabled")
	}

	log.Fatal(router.Run(":" + cfg.Port))
}
