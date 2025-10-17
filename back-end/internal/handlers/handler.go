package handlers

import "github.com/gin-gonic/gin"

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Set up static file serving
	h.SetupStaticFileMiddleware(router)
	h.SetupSPAFallback(router)

	// API routes
	router.GET("/api/health", h.Health)

	return router
}
