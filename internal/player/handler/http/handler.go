package http

import (
	"net/http"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/player"
	httpErr "github.com/Unlites/nba_api/pkg/http_errors"
	"github.com/Unlites/nba_api/pkg/utils"
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

	if err := c.BindJSON(player); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.playerUC.Create(player); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *playerHandler) GetById(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	player, err := h.playerUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, player)
}

func (h *playerHandler) Update(c *gin.Context) {
	player := &models.Player{}

	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	player.Id = idInput

	if err := c.BindJSON(player); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.playerUC.Update(player); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *playerHandler) Delete(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.playerUC.Delete(idInput); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
