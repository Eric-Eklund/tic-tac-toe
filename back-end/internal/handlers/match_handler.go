package handlers

import (
	"net/http"
	"tic-tac-toe/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NewMatch(c *gin.Context) {
	match := models.NewMatch()

	c.JSON(http.StatusOK, match)
}
