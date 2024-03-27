package router

import "github.com/gin-gonic/gin"

// Por ter "I" maiusculo ele é automaticamente exportado. Minusculo é variavel local.
func Initialize() {
	//Inicializa o route utilizando as configurações Default do gin
	router := gin.Default()
	//Definindo uma rota
	initializeRoutes(router)
	//Rodando a API
	router.Run(":8081")
}
