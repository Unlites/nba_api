package http

import (
	"github.com/Unlites/nba_api/internal/stat"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(statGroup *gin.RouterGroup, h stat.Handler) {
	statGroup.POST("/", h.Create)
	statGroup.GET("/:id", h.GetById)
}
