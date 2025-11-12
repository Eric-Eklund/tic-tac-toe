package handlers

import (
	"os"
	"path/filepath"
	"tic-tac-toe/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(config.GetCors()))
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// API routes (before static files)
	router.GET("/api/health", h.Health)
	router.GET("/api/new-match", h.NewMatch)

	// Serve React static assets first (more specific routes)
	router.Static("/react/assets", "../front-end/dist/assets")
	router.StaticFile("/react/vite.svg", "../front-end/dist/vite.svg")

	// Serve Godot game static files
	router.Static("/game-assets", "../game-client/build")

	// Root route serves Godot game
	router.GET("/", func(c *gin.Context) {
		c.File("../game-client/build/index.html")
	})

	// React app at /react
	router.GET("/react", func(c *gin.Context) {
		c.File("../front-end/dist/index.html")
	})

	// Fallback for any unmatched routes
	router.NoRoute(func(c *gin.Context) {
		path := c.Request.URL.Path

		// If path starts with /react, serve React app
		if len(path) >= 6 && path[:6] == "/react" {
			c.File("../front-end/dist/index.html")
			return
		}

		// Try to serve from Godot build folder (exact file)
		filePath := filepath.Join("../game-client/build", path)
		if fileExists(filePath) {
			c.File(filePath)
			return
		}

		// Only serve index.html for HTML navigation (not for .js, .wasm, etc.)
		if filepath.Ext(path) == "" || filepath.Ext(path) == ".html" {
			c.File("../game-client/build/index.html")
			return
		}

		// For everything else, 404
		c.Status(404)
	})

	return router
}

// Helper function to check if file exists
func fileExists(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}