package http

import (
	"net/http"

	"github.com/Unlites/nba_api/internal/game"
	"github.com/Unlites/nba_api/internal/models"
	httpErr "github.com/Unlites/nba_api/pkg/http_errors"
	"github.com/Unlites/nba_api/pkg/utils"
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
	if err := c.BindJSON(game); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.gameUC.Create(game); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *gameHandler) GetById(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	game, err := h.gameUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, game)
}

func (h *gameHandler) Update(c *gin.Context) {
	game := &models.Game{}

	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	game.Id = idInput

	if err := c.BindJSON(game); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.gameUC.Update(game); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *gameHandler) Delete(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.gameUC.Delete(idInput); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
