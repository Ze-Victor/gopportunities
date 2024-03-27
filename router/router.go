package router

import "github.com/gin-gonic/gin"

// Por ter "I" maiusculo ele é automaticamente exportado. Minusculo é variavel local.
func Initialize() {
	//Inicializa o route utilizando as configurações Default do gin
	router := gin.Default()
	//Definindo uma rota
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	//Rodando a API
	router.Run(":8081")
}
