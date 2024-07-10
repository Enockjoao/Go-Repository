package router

import "github.com/gin-gonic/gin"

func Initialize() {
	//Inicializando o Router utilizando as configurações Defaut do gin
	Router := gin.Default()
	//Definindo uma Rota
	Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Estamos Rodando a nossa Api
	Router.Run(":8080")
}
