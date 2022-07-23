package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nelsonomoi/usingJWT/controllers"
	"github.com/nelsonomoi/usingJWT/middlewares"
)

func InitRoutes() *gin.Engine {

	router := gin.Default()
	api := router.Group("/api")
	{
		api.POST("/token",controllers.GenerateToken)
		api.POST("/user/register",controllers.RegisterUser)
		secured := api.Group("/secured").Use(middlewares.Auth())
		{
			secured.GET("/ping",controllers.Ping)
		}
	}

	return router
}