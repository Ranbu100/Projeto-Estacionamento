package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListEntradasSaidasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Listando entradas/saidas de veiculos",
	})
}
func ListPagamentosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Listando pagamentos",
	})
}
func ListUsuariosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Listando usuarios",
	})
}
func ListVagasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Listando vagas de estacionamento",
	})
}
func ListVeiculosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Listando veiculos",
	})
}
func ListReservasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Listando reservas de vagas",
	})
}
