package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":    "ok",
		"service":   "tic-tac-toe-backend",
		"version":   "1.0.0",
		"timestamp": time.Now().Format(time.RFC3339),
	})
}
