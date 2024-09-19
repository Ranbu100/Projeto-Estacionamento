package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteEntradasSaidasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Excluindo entrada/saida de veiculos",
	})
}
func DeletePagamentosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Excluindo pagamento",
	})
}
func DeleteReservasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Excluindo reservas de estacionamento",
	})
}
func DeleteUsuariosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Excluindo usuario",
	})
}
func DeleteVeiculosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Excluindo veiculo",
	})
}
func DeleteVagasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Excluindo vaga de estacionamento",
	})
}
