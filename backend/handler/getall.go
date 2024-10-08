package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/enums"
	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/gin-gonic/gin"
)

func ListEntradasSaidasHandler(ctx *gin.Context) {
	var entradasSaidas []schemas.EntradasSaidas

	// Buscar todas as entradas/saídas no banco de dados
	if err := db.Find(&entradasSaidas).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar entradas/saídas"})
		return
	}

	// Retornar a lista de entradas/saídas em JSON
	ctx.JSON(http.StatusOK, gin.H{"entradas_saidas": entradasSaidas})
}

func ListPagamentosHandler(ctx *gin.Context) {
	var pagamentos []schemas.Pagamentos

	// Buscar todos os pagamentos no banco de dados
	if err := db.Find(&pagamentos).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar pagamentos"})
		return
	}

	// Retornar a lista de pagamentos em JSON
	ctx.JSON(http.StatusOK, gin.H{"pagamentos": pagamentos})
}

func ListUsuariosHandler(ctx *gin.Context) {
	var usuarios []schemas.Usuarios

	// Buscar todos os usuários no banco de dados
	if err := db.Find(&usuarios).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar usuários"})
		return
	}

	// Retornar a lista de usuários em JSON
	ctx.JSON(http.StatusOK, gin.H{"usuarios": usuarios})
}

func mapStatusVaga(status int) string {
	switch status {
	case enums.Disponivel:
		return "Disponível"
	case enums.Ocupada:
		return "Ocupada"
	case enums.Reservada:
		return "Reservada"
	case enums.Manutencao:
		return "Manutenção"
	default:
		return "Desconhecido"
	}
}

func mapTipoVaga(tipo uint8) string {
	switch tipo {
	case enums.TipoCarro:
		return "Carro"
	case enums.TipoMoto:
		return "Moto"
	default:
		return "Desconhecido"
	}
}

func ListVagasHandler(ctx *gin.Context) {
	var vagas []schemas.Vagas

	// Buscar todas as vagas no banco de dados
	if err := db.Find(&vagas).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar vagas"})
		return
	}

	var response []struct {
		NumeroVaga int    `json:"numero_vaga"`
		TipoVaga   string `json:"tipo_vaga"`
		Status     string `json:"status_vaga"`
	}

	for _, vaga := range vagas {
		response = append(response, struct {
			NumeroVaga int    `json:"numero_vaga"`
			TipoVaga   string `json:"tipo_vaga"`
			Status     string `json:"status_vaga"`
		}{
			NumeroVaga: vaga.NumeroVaga,
			TipoVaga:   mapTipoVaga(vaga.TipoVaga), // Convertendo para texto
			Status:     mapStatusVaga(vaga.Status), // Convertendo para texto
		})
	}

	// Retornar o JSON corretamente
	ctx.JSON(http.StatusOK, response)
}

func ListVeiculosHandler(ctx *gin.Context) {
	var veiculos []schemas.Veiculos

	// Buscar todos os veículos no banco de dados
	if err := db.Find(&veiculos).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar veículos"})
		return
	}

	// Retornar a lista de veículos em JSON
	ctx.JSON(http.StatusOK, gin.H{"veiculos": veiculos})
}

func ListReservasHandler(ctx *gin.Context) {
	var reservas []schemas.Reservas

	// Buscar todas as reservas no banco de dados
	if err := db.Find(&reservas).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao listar reservas"})
		return
	}

	// Retornar a lista de reservas em JSON
	ctx.JSON(http.StatusOK, gin.H{"reservas": reservas})
}
