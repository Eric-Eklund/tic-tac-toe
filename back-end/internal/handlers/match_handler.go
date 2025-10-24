package handlers

import (
	"net/http"
	"tic-tac-toe/internal/models"
	"tic-tac-toe/internal/services"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NewMatch(c *gin.Context) {
	match := models.NewMatch()

	services.SaveMatch(match)

	c.JSON(http.StatusOK, match)
}
