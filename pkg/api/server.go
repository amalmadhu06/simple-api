package http

import (
	"github.com/amalmadhu06/simple-api/pkg/api/handler"
	"github.com/gin-gonic/gin"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler) *ServerHTTP {
	engine := gin.Default()

	api := engine.Group("/api", gin.Recovery())
	api.GET("users", userHandler.FindAll)

	return &ServerHTTP{
		engine: engine,
	}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
