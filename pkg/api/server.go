package http

import (
	_ "github.com/abhinandpn/project-ecom/cmd/api/docs"
	"github.com/abhinandpn/project-ecom/pkg/api/handler"
	"github.com/abhinandpn/project-ecom/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler *handler.UserHandler,
	adminHandler *handler.AdminHandler) *ServerHTTP {

	Engine := gin.New()
	Engine.Use(gin.Logger())
	// Engine.LoadHTMLGlob("*.html")

	// For Swagger
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// For Routes
	routes.UserRoutes(Engine.Group("/"), userHandler)
	routes.AdminRoute(Engine.Group("/admin"), adminHandler)

	return &ServerHTTP{engine: Engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
