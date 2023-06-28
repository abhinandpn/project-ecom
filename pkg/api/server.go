package http

import (
	_ "github.com/abhinandpn/project-ecom/cmd/api/docs"
	handlerInterface "github.com/abhinandpn/project-ecom/pkg/api/handler/interfaces"
	"github.com/abhinandpn/project-ecom/pkg/api/routes"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(userHandler handlerInterface.UserHandler,
	adminHandler handlerInterface.AdminHandler,
	productHandler handlerInterface.ProductHandler,
	cartHandler handlerInterface.CartHandler,
	orderHandler handlerInterface.OrderHandler) *ServerHTTP {

	Engine := gin.New()
	Engine.Use(gin.Logger())

	// For Html
	// Engine.LoadHTMLGlob("*.html")

	// For Swagger
	Engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// For Routes
	routes.UserRoutes(Engine.Group("/"), userHandler, productHandler, cartHandler, orderHandler)
	routes.AdminRoute(Engine.Group("/admin"), adminHandler, productHandler, orderHandler)

	return &ServerHTTP{engine: Engine}
}

func (sh *ServerHTTP) Start() {
	sh.engine.Run(":3000")
}
