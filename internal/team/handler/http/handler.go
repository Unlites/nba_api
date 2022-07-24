package http

import (
	"net/http"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/team"
	httpErr "github.com/Unlites/nba_api/pkg/http_errors"
	"github.com/Unlites/nba_api/pkg/utils"
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

	if err := c.BindJSON(team); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.teamUC.Create(team); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *teamHandler) GetById(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	team, err := h.teamUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, team)
}

func (h *teamHandler) Update(c *gin.Context) {
	team := &models.Team{}

	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	team.Id = idInput

	if err := c.BindJSON(team); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.teamUC.Update(team); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}
}

func (h *teamHandler) Delete(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.teamUC.Delete(idInput); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}
}
