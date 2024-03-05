package router

import (
	"github.com/gin-gonic/gin"
	"github.com/zeimedee/saber/internal/handlers"
	"github.com/zeimedee/saber/internal/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	value := services.NewValue()
	valueHandler := handlers.NewValueHandler(value)

	api := router.Group("/saber")
	{
		api.GET("/", handlers.HealthCheck)
		api.POST("/add", valueHandler.AddValue)
		api.GET("/ws", handlers.WebSocket)
	}
	return router
}
