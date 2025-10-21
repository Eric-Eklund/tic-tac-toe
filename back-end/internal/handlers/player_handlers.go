package handlers

import (
	"net/http"
	"tic-tac-toe/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) InitialPlayers(c *gin.Context) {
	players, err := models.GetInitialPlayers()
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, players)
}
