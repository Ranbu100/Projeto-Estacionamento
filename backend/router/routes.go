package router

import (
	"time"

	"github.com/Ranbu100/Projeto-Estacionamento/handler"
	"github.com/Ranbu100/Projeto-Estacionamento/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(router *gin.Engine) {
	// Configurando o middleware CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Permitir todas as origens, mas é recomendável especificar
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Iniciando o handler
	handler.InitializeHandler()

	v1 := router.Group("/api/v1")
	{
		// Rotas públicas (sem proteção JWT)
		v1.POST("/login", handler.LoginHandler) // Login e geração do token JWT
		v1.POST("/usuario", handler.CreateUsuariosHandler)
		// Rotas protegidas com JWT
		v1.Use(middleware.AuthMiddleware())
		{
			// Rotas para as vagas
			v1.GET("/vagas", handler.ListVagasHandler)
			v1.GET("/vagas/:id", handler.ShowVagasHandler)
			v1.POST("/vagas", handler.CreateVagasHandler)
			v1.PUT("/vagas/:id", handler.UpdateVagasHandler)
			v1.DELETE("/vagas/:id", handler.DeleteVagasHandler)

			// Rotas para usuários
			v1.GET("/usuarios", handler.ListUsuariosHandler)
			v1.GET("/usuarios/:id", handler.ShowUsuariosHandler)
			v1.POST("/usuarios", handler.CreateUsuariosHandler)
			v1.PUT("/usuarios/:id", handler.UpdateUsuariosHandler)
			v1.DELETE("/usuarios/:id", handler.DeleteUsuariosHandler)

			// Rotas para veículos
			v1.GET("/veiculos", handler.ListVeiculosHandler)
			v1.GET("/veiculos/:id", handler.ShowVeiculosHandler)
			v1.POST("/veiculos", handler.CreateVeiculosHandler)
			v1.PUT("/veiculos/:id", handler.UpdateVeiculosHandler)
			v1.DELETE("/veiculos/:id", handler.DeleteVeiculosHandler)

			// Rotas para entradas e saídas
			v1.GET("/entradas-saidas", handler.ListEntradasSaidasHandler)
			v1.GET("/entradas-saidas/:id", handler.ShowEntradasSaidasHandler)
			v1.POST("/entradas-saidas", handler.CreateEntradasSaidasHandler)
			v1.PUT("/entradas-saidas/:id", handler.UpdateEntradasSaidasHandler)
			v1.DELETE("/entradas-saidas/:id", handler.DeleteEntradasSaidasHandler)

			// Rotas para pagamentos
			v1.GET("/pagamentos", handler.ListPagamentosHandler)
			v1.GET("/pagamentos/:id", handler.ShowPagamentosHandler)
			v1.POST("/pagamentos", handler.CreatePagamentosHandler)
			v1.PUT("/pagamentos/:id", handler.UpdatePagamentosHandler)
			v1.DELETE("/pagamentos/:id", handler.DeletePagamentosHandler)

			// Rotas para reservas
			v1.GET("/reservas", handler.ListReservasHandler)
			v1.GET("/reservas/:id", handler.ShowReservasHandler)
			v1.POST("/reservas", handler.CreateReservasHandler)
			v1.PUT("/reservas/:id", handler.UpdateReservasHandler)
			v1.DELETE("/reservas/:id", handler.DeleteReservasHandler)
		}
	}
}
