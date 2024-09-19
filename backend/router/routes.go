package router

import (
	"github.com/Ranbu100/Projeto-Estacionamento/handler"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	//inciando o handler
	handler.InitializeHandler()
	v1 := router.Group("/api/v1")
	{
		//rotas para as vagas
		v1.GET("/vagas", handler.ListVagasHandler)
		v1.GET("/vagas/:id", handler.ShowVagasHandler)
		v1.POST("/vagas", handler.CreateVagasHandler)
		v1.PUT("/vagas/:id", handler.UpdateVagasHandler)
		v1.DELETE("/vagas/:id", handler.DeleteVagasHandler)

		//rotas para usuarios
		v1.GET("/usuarios", handler.ListUsuariosHandler)
		v1.GET("/usuarios/:id", handler.ShowUsuariosHandler)
		v1.POST("/usuarios", handler.CreateUsuariosHandler)
		v1.PUT("/usuarios/:id", handler.UpdateUsuariosHandler)
		v1.DELETE("/usuarios/:id", handler.DeleteUsuariosHandler)

		//rotas para veiculos
		v1.GET("/veiculos", handler.ListVeiculosHandler)
		v1.GET("/veiculos/:id", handler.ShowVeiculosHandler)
		v1.POST("/veiculos", handler.CreateVeiculosHandler)
		v1.PUT("/veiculos/:id", handler.UpdateVeiculosHandler)
		v1.DELETE("/veiculos/:id", handler.DeleteVeiculosHandler)

		//rotas para entradas e saidas
		v1.GET("/entradas-saidas", handler.ListEntradasSaidasHandler)
		v1.GET("/entradas-saidas/:id", handler.ShowEntradasSaidasHandler)
		v1.POST("/entradas-saidas", handler.CreateEntradasSaidasHandler)
		v1.PUT("/entradas-saidas/:id", handler.UpdateEntradasSaidasHandler)
		v1.DELETE("/entradas-saidas/:id", handler.DeleteEntradasSaidasHandler)

		//rotas para pagamentos
		v1.GET("/pagamentos", handler.ListPagamentosHandler)
		v1.GET("/pagamentos/:id", handler.ShowPagamentosHandler)
		v1.POST("/pagamentos", handler.CreatePagamentosHandler)
		v1.PUT("/pagamentos/:id", handler.UpdatePagamentosHandler)
		v1.DELETE("/pagamentos/:id", handler.DeletePagamentosHandler)

		//rotas para reservas
		v1.GET("/reservas", handler.ListReservasHandler)
		v1.GET("/reservas/:id", handler.ShowReservasHandler)
		v1.POST("/reservas", handler.CreateReservasHandler)
		v1.PUT("/reservas/:id", handler.UpdateReservasHandler)
		v1.DELETE("/reservas/:id", handler.DeleteReservasHandler)
	}
}
