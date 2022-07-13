package http

import (
	"net/http"
	"strconv"

	"github.com/Unlites/nba_api/internal/models"
	"github.com/Unlites/nba_api/internal/stat"
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

	if err := h.statUC.Create(stat); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusOK)
}

func (h *statHandler) GetById(c *gin.Context) {
	idInput, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	stat, err := h.statUC.GetById(idInput)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, stat)
}
