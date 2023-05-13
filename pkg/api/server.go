package http

import (
	"github.com/gin-gonic/gin"

	_ "github.com/abhinandpn/project-ecom/cmd/api/docs"
	"github.com/abhinandpn/project-ecom/pkg/api/handler"
	"github.com/abhinandpn/project-ecom/pkg/api/routes"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler) *ServerHTTP {

	Engine := gin.New()
	Engine.Use(gin.Logger())

	// For Swagger
	// For Routes
	routes.UserRoutes(Engine.Group("/"), userHandler)
	routes.AdminRoute(Engine.Group("/admin"), adminHandler)

	return &ServerHTTP{engine: Engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
