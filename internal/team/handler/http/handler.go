package http

import (
	"net/http"
	"strconv"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/team"
	"github.com/gin-gonic/gin"
)

type teamHandler struct {
	teamUC team.UseCase
}

func NewTeamHandler(teamUC team.UseCase) team.Handler {
	return &teamHandler{teamUC: teamUC}
}

func (h *teamHandler) Create(c *gin.Context) {
	team := &models.Team{}

	if err := h.teamUC.Create(team); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
	}

	c.Status(http.StatusOK)
}

func (h *teamHandler) GetById(c *gin.Context) {
	idInput, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	team, err := h.teamUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, team)
}
