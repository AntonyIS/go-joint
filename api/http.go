package api

import (
	"github.com/gin-gonic/gin"
)

type RouteHandler interface {
	Get(*gin.Context)
	GetAll(*gin.Context)
	Post(*gin.Context)
	Put(*gin.Context)
	Delete(*gin.Context)
}
