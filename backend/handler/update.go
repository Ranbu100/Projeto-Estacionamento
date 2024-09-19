package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateEntradasSaidasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando entradas/saidas de veiculos",
	})
}
func UpdatePagamentosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando pagamentos",
	})
}
func UpdateReservasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando reservas",
	})
}
func UpdateUsuariosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando usuario",
	})
}
func UpdateVeiculosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando veiculo",
	})
}
func UpdateVagasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando vagas de estacionamento",
	})
}
