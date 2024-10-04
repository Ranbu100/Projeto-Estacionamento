package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/Ranbu100/Projeto-Estacionamento/utils"
	"github.com/gin-gonic/gin"
)

func UpdateEntradasSaidasHandler(ctx *gin.Context) {
	var input struct {
		VeiculoID uint   `json:"veiculo_id"`
		VagaID    uint   `json:"vaga_id"`
		Entrada   string `json:"entrada"`
		Saida     string `json:"saida"`
	}

	// Bind JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Parse entrada (obrigatório)
	entrada, err := utils.ParseTime(input.Entrada)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de entrada inválido"})
		return
	}

	// Parse saída (opcional)
	saida, err := utils.ParseOptionalTime(&input.Saida)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de saída inválido"})
		return
	}

	// Obter o ID da entrada/saída
	id := ctx.Param("id")

	// Procurar a entrada/saída no banco de dados
	var entradaSaida schemas.EntradasSaidas
	if err := db.First(&entradaSaida, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Entrada/Saída não encontrada"})
		return
	}

	// Atualizar os dados
	entradaSaida.VeiculoId = input.VeiculoID
	entradaSaida.VagaId = input.VagaID
	entradaSaida.Entrada = entrada
	entradaSaida.Saida = saida

	// Salvar as alterações
	if err := db.Save(&entradaSaida).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a entrada/saída"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Entrada/Saída atualizada com sucesso"})
}

func UpdatePagamentosHandler(ctx *gin.Context) {
	var input struct {
		Valor  float64 `json:"valor"`
		Metodo string  `json:"metodo"`
		Status string  `json:"status"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	id := ctx.Param("id")

	var pagamento schemas.Pagamentos
	if err := db.First(&pagamento, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Pagamento não encontrado"})
		return
	}

	pagamento.Valor = input.Valor
	pagamento.Metodo = input.Metodo
	pagamento.Status = input.Status

	if err := db.Save(&pagamento).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar pagamento"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Pagamento atualizado com sucesso"})
}

func UpdateReservasHandler(ctx *gin.Context) {
	var input struct {
		UsuarioID uint   `json:"usuario_id"`
		VagaID    uint   `json:"vaga_id"`
		Data      string `json:"data"`
		Status    string `json:"status"`
	}

	// Bind JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Parse data (obrigatório)
	data, err := utils.ParseTime(input.Data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato de data inválido"})
		return
	}

	// Obter o ID da reserva
	id := ctx.Param("id")

	// Procurar a reserva no banco de dados
	var reserva schemas.Reservas
	if err := db.First(&reserva, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Reserva não encontrada"})
		return
	}

	// Atualizar os dados
	reserva.UsuarioId = input.UsuarioID
	reserva.VagaId = input.VagaID
	reserva.DataHora = data
	reserva.Status = input.Status

	// Salvar as alterações
	if err := db.Save(&reserva).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a reserva"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Reserva atualizada com sucesso"})
}

func UpdateUsuariosHandler(ctx *gin.Context) {
	var input struct {
		Nome  string `json:"nome"`
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	// Bind JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Obter o ID do usuário da URL
	id := ctx.Param("id")

	// Buscar o usuário no banco de dados
	var usuario schemas.Usuarios
	if err := db.First(&usuario, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Usuário não encontrado"})
		return
	}

	// Atualizar os campos
	usuario.Nome = input.Nome
	usuario.Email = input.Email
	hashedPassword, err := utils.HashPassword(input.Senha)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}
	usuario.Senha = hashedPassword

	// Salvar as alterações
	if err := db.Save(&usuario).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar usuário"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Usuário atualizado com sucesso"})
}

func UpdateVeiculosHandler(ctx *gin.Context) {
	var input struct {
		Placa     string `json:"placa"`
		Modelo    string `json:"modelo"`
		Cor       string `json:"cor"`
		UsuarioID uint   `json:"usuario_id"`
	}

	// Bind JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Obter o ID do veículo
	id := ctx.Param("id")

	// Procurar o veículo no banco de dados
	var veiculo schemas.Veiculos
	if err := db.First(&veiculo, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Veículo não encontrado"})
		return
	}

	// Atualizar os dados
	veiculo.Placa = input.Placa
	veiculo.Modelo = input.Modelo
	veiculo.Cor = input.Cor
	veiculo.UsuarioId = input.UsuarioID

	// Salvar as alterações
	if err := db.Save(&veiculo).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar o veículo"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Veículo atualizado com sucesso"})
}

func UpdateVagasHandler(ctx *gin.Context) {
	var input struct {
		NumeroVaga int   `json:"numero_vaga"`
		TipoVaga   uint8 `json:"tipo_vaga"`
		Status     int   `json:"status_vaga"`
	}

	// Bind JSON
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Obter o ID da vaga
	id := ctx.Param("id")

	// Procurar a vaga no banco de dados
	var vaga schemas.Vagas
	if err := db.First(&vaga, id).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Vaga não encontrada"})
		return
	}

	// Atualizar os dados
	vaga.NumeroVaga = input.NumeroVaga
	vaga.TipoVaga = input.TipoVaga
	vaga.Status = input.Status

	// Salvar as alterações
	if err := db.Save(&vaga).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar a vaga"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Vaga atualizada com sucesso"})
}
