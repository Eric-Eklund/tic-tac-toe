package handlers

import (
	"net/http"
	"tic-tac-toe/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitialPlayers(c *gin.Context) {
	players := models.GetInitialPlayers()

	c.JSON(http.StatusOK, players)
}
