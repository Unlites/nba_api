package player

import "github.com/gin-gonic/gin"

type Handler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}
