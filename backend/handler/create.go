package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateEntradasSaidasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando entrada/saidas de veiculos",
	})
}
func CreatePagamentosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando pagamentos",
	})
}
func CreateReservasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando reservas de vagas",
	})
}
func CreateUsuariosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando usuario",
	})
}
func CreateVagasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando vagas de estacionamento",
	})
}
func CreateVeiculosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando veiculos",
	})
}
