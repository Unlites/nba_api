package http

import (
	"net/http"
	"strconv"

	"github.com/Unlites/nba_api/internal/game"
	"github.com/Unlites/nba_api/internal/models"
	"github.com/gin-gonic/gin"
)

type gameHandler struct {
	gameUC game.UseCase
}

func NewGameHandler(gameUC game.UseCase) game.Handler {
	return &gameHandler{gameUC: gameUC}
}

func (h *gameHandler) Create(c *gin.Context) {
	game := &models.Game{}

	if err := h.gameUC.Create(game); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusOK)
}

func (h *gameHandler) GetById(c *gin.Context) {
	idInput, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	game, err := h.gameUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, game)
}
