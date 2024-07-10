package router

import "github.com/gin-gonic/gin"

func Initialize() {

	//Initialize Router
	router := gin.Default()

	//Initialize Routers
	initializeRoutes(router)

	//Estamos Rodando a nossa Api
	router.Run(":8080")
}
