package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/gin-gonic/gin"
)

func ShowEntradasSaidasHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var entradaSaida schemas.EntradasSaidas

	if err := db.First(&entradaSaida, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Entrada/Saída não encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, entradaSaida)
}

func ShowPagamentosHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var pagamento schemas.Pagamentos

	if err := db.First(&pagamento, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Pagamento não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, pagamento)
}

func ShowReservasHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var reserva schemas.Reservas

	if err := db.First(&reserva, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Reserva não encontrada"})
		return
	}

	ctx.JSON(http.StatusOK, reserva)
}

func ShowUsuariosHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var usuario schemas.Usuarios

	if err := db.First(&usuario, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, usuario)
}

func ShowVeiculosHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var veiculo schemas.Veiculos

	if err := db.First(&veiculo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Veiculo nao encontrado"})
		return
	}
}
func ShowVagasHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	var vaga schemas.Vagas

	if err := db.First(&vaga, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Vaga não encontrada"})
	}
}
