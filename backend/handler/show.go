package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowEntradasSaidasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Exibindo entradas/saidas de veiculos",
	})
}
func ShowPagamentosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Exibindo pagamentos",
	})
}
func ShowReservasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Exibindo reservas de estacionamentos",
	})
}
func ShowUsuariosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Exibindo usuarios",
	})
}
func ShowVeiculosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Exibindo veiculos",
	})
}
func ShowVagasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Exibindo vagas de estacionamento",
	})
}
