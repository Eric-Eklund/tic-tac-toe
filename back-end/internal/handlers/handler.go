package handlers

import (
	"tic-tac-toe/internal/config"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)
import _ "github.com/gin-contrib/cors"

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

	// Set up static file serving
	h.SetupStaticFileMiddleware(router)
	h.SetupSPAFallback(router)

	// API routes
	router.GET("/api/health", h.Health)
	router.GET("/api/new-match", h.NewMatch)

	return router
}
