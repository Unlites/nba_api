package http

import (
	"github.com/Unlites/nba_api/internal/game"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(gameGroup *gin.RouterGroup, h game.Handler) {
	gameGroup.POST("/", h.Create)
	gameGroup.GET("/:id", h.GetById)
}
