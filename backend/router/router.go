package router

import (
	"github.com/Ranbu100/Projeto-Estacionamento/handler"
	"github.com/gin-gonic/gin"
)

func Initialize() {
	//incializa o roteamento usando as configuraçoes padroes do gin
	router := gin.Default()

	//incilizando as rotas
	InitializeRoutes(router)

	handler.InitializeHandler()

	//inciando servidor
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
