package handlers

import (
	"net/http"
	"tic-tac-toe/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GameBoard(c *gin.Context) {
	gameBoard := models.NewGameBoard()

	c.JSON(http.StatusOK, gameBoard)
}
