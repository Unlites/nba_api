package http

import (
	"net/http"
	"strconv"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/player"
	"github.com/gin-gonic/gin"
)

type playerHandler struct {
	playerUC player.UseCase
}

func NewPlayerHandler(playerUC player.UseCase) player.Handler {
	return &playerHandler{playerUC: playerUC}
}

func (h *playerHandler) Create(c *gin.Context) {
	player := &models.Player{}

	if err := h.playerUC.Create(player); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusOK)
}

func (h *playerHandler) GetById(c *gin.Context) {
	idInput, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	player, err := h.playerUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, player)
}
