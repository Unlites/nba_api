package http

import (
	"github.com/Unlites/nba_api/internal/player"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(playerGroup *gin.RouterGroup, h player.Handler) {
	playerGroup.POST("/", h.Create)
	playerGroup.GET("/:id", h.GetById)
}
