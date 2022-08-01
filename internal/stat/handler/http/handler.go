package http

import (
	"net/http"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/stat"
	httpErr "github.com/Unlites/nba_api/pkg/http_errors"
	"github.com/Unlites/nba_api/pkg/utils"
	"github.com/gin-gonic/gin"
)

type statHandler struct {
	statUC stat.UseCase
}

func NewStatHandler(statUC stat.UseCase) stat.Handler {
	return &statHandler{statUC: statUC}
}

func (h *statHandler) Create(c *gin.Context) {
	stat := &models.Stat{}

	if err := c.BindJSON(stat); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.statUC.Create(stat); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *statHandler) GetById(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	stat, err := h.statUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, stat)
}

func (h *statHandler) GetAvgByPlayerId(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("player_id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	avgStat, err := h.statUC.GetAvgByPlayerId(idInput)
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.JSON(http.StatusOK, avgStat)
}

func (h *statHandler) Update(c *gin.Context) {
	stat := &models.Stat{}

	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	stat.Id = idInput

	if err := c.BindJSON(stat); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.statUC.Update(stat); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}

func (h *statHandler) Delete(c *gin.Context) {
	idInput, err := utils.ParseId(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	if err := h.statUC.Delete(idInput); err != nil {
		c.AbortWithStatusJSON(httpErr.NewErrorResponse(err.Error()))
		return
	}

	c.Status(http.StatusOK)
}
