package http

import (
	"github.com/Unlites/nba_api/internal/team"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(teamGroup *gin.RouterGroup, h team.Handler) {
	teamGroup.POST("/", h.Create)
	teamGroup.GET("/:id", h.GetById)
	teamGroup.PUT("/:id", h.Update)
	teamGroup.DELETE("/:id", h.Delete)
}
