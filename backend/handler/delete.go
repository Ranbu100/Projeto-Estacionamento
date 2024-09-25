package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/gin-gonic/gin"
)

func DeleteEntradasSaidasHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := db.Delete(&schemas.EntradasSaidas{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir entrada/saída"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Entrada/Saída excluída com sucesso"})
}

func DeletePagamentosHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := db.Delete(&schemas.Pagamentos{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir pagamento"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pagamento excluído com sucesso"})
}

func DeleteReservasHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := db.Delete(&schemas.Reservas{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir reserva"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva excluída com sucesso"})
}

func DeleteUsuariosHandler(ctx *gin.Context) {
	// Pegando o ID do usuário a ser deletado via URL
	id := ctx.Param("id")

	// Tentando deletar o usuário no banco de dados
	if err := db.Delete(&schemas.Usuarios{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuário excluído com sucesso"})
}

func DeleteVeiculosHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := db.Delete(&schemas.Veiculos{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir veículo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Veículo excluído com sucesso"})
}

func DeleteVagasHandler(ctx *gin.Context) {
	id := ctx.Param("id")

	if err := db.Delete(&schemas.Vagas{}, id).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao excluir vaga"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vaga excluída com sucesso"})
}
