package handler

import (
	"net/http"
	"os"
	"time"

	"github.com/Ranbu100/Projeto-Estacionamento/middleware"
	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/Ranbu100/Projeto-Estacionamento/utils"
	"github.com/gin-gonic/gin"
)

func CreateEntradasSaidasHandler(ctx *gin.Context) {
	var input struct {
		VeiculoID uint    `json:"veiculo_id"`
		VagaID    uint    `json:"vaga_id"`
		Entrada   string  `json:"entrada"`
		Saida     *string `json:"saida,omitempty"`
	}

	// Bind JSON para a struct de input
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	entradaSaida := schemas.EntradasSaidas{
		VeiculoId: input.VeiculoID,
		VagaId:    input.VagaID,
	}

	// Parse da data/hora de entrada
	entrada, err := utils.ParseTime(input.Entrada)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido para a data de entrada"})
		return
	}
	entradaSaida.Entrada = entrada

	// Parse da data/hora de saída (opcional)
	saida, err := utils.ParseOptionalTime(input.Saida)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Formato inválido para a data de saída"})
		return
	}
	entradaSaida.Saida = saida

	// Criar entrada/saída no banco de dados
	if err := db.Create(&entradaSaida).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar no banco de dados"})
		return
	}

	logger.Info("Entrada/saída criada com sucesso")
	ctx.JSON(http.StatusCreated, gin.H{"message": "Entrada/saída criada com sucesso"})
}

func CreatePagamentosHandler(ctx *gin.Context) {
	var input struct {
		EntradasSaidasId uint    `json:"entradas-saidas_id"`
		Valor            float64 `json:"valor_pago"`
		Metodo           string  `json:"metodo_pagamento"`
		Status           string  `json:"status_pagamento"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}
	pagamentos := schemas.Pagamentos{
		EntradasSaidasId: input.EntradasSaidasId,
		Valor:            input.Valor,
		Metodo:           input.Metodo,
		Status:           input.Status,
	}

	if err := db.Create(&pagamentos).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar pagamento no banco de dados"})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Pagamento criado com sucesso"})
}
func CreateReservasHandler(c *gin.Context) {
	var input struct {
		VagaId    uint      `json:"vaga_id"`
		UsuarioId uint      `json:"usuario_id"`
		DataHora  time.Time `gorm:"not null" json:"data_hora_reserva"`
		Duracao   uint      `gorm:"not null" json:"duracao"`
		Status    string    `gorm:"size:50;not null" json:"status_reserva"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Criar nova reserva
	reserva := schemas.Reservas{
		VagaId:    input.VagaId,
		UsuarioId: input.UsuarioId,
		DataHora:  input.DataHora,
		Duracao:   input.Duracao,
		Status:    input.Status,
	}

	if err := db.Create(&reserva).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar reserva no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Reserva criada com sucesso"})
}

// @Summary Cria um usuário
// @Description Cria um novo usuário no sistema
// @Tags usuários
// @Accept  json
// @Produce  json
// @Param   usuario  body  handler.CreateUsuarioInput  true  "Dados do usuário"
// @Success 201 {object} gin.H{"message": "Usuário criado com sucesso"}
// @Failure 400 {object} gin.H{"error": "Dados inválidos"}
// @Router /usuarios [post]
func CreateUsuariosHandler(c *gin.Context) {
	var input struct {
		Nome     string `json:"nome"`
		Email    string `json:"email"`
		Senha    string `json:"senha"`
		Telefone string `json:"telefone"`
	}

	// Bind JSON para a struct de input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verificar se o email já existe
	var usuarioExistente schemas.Usuarios
	if err := db.Where("email = ?", input.Email).First(&usuarioExistente).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Email já cadastrado"})
		return
	}

	// Gerar o hash da senha
	hashedPassword, err := utils.HashPassword(input.Senha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}

	// Criar o novo usuário
	usuario := schemas.Usuarios{
		Nome:     input.Nome,
		Email:    input.Email,
		Senha:    hashedPassword,
		Telefone: input.Telefone,
	}
	if err := db.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar usuário no banco de dados"})
		return
	}
	logger.Info("Usuário criado com sucesso: ", input.Email)
	c.JSON(http.StatusCreated, gin.H{"message": "Usuário criado com sucesso"})
}

func LoginHandler(ctx *gin.Context) {
	var input struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	// JSON recebido
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verificar se as credenciais são de um admin
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminSenha := os.Getenv("ADMIN_SENHA")

	// Caso seja admin, checar as credenciais diretamente
	if input.Email == adminEmail && input.Senha == adminSenha {
		// Gerar token JWT para o admin
		token, err := middleware.GenerateJWT("admin")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "Login de admin bem-sucedido!"})
		return
	}

	// Caso contrário, continue verificando no banco de dados para usuários comuns
	var usuario schemas.Usuarios
	if err := db.Where("email = ?", input.Email).First(&usuario).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Comparar senhas
	if !utils.CheckPasswordHash(input.Senha, usuario.Senha) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gerar token JWT para o usuário comum
	token, err := middleware.GenerateJWT(usuario.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	// Se a senha estiver correta
	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "Login bem-sucedido"})
}

func CreateVagasHandler(c *gin.Context) {
	var input struct {
		NumeroVaga int    `json:"numero_vaga"`
		TipoVaga   string `json:"tipo_vaga"`
		Status     string `json:"status_vaga"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Criar nova vaga
	vaga := schemas.Vagas{
		NumeroVaga: input.NumeroVaga,
		TipoVaga:   input.TipoVaga,
		Status:     input.Status,
	}

	if err := db.Create(&vaga).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar vaga no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Vaga criada com sucesso"})
}

func CreateVeiculosHandler(c *gin.Context) {
	var input struct {
		Placa     string `json:"placa"`
		Modelo    string `json:"modelo"`
		Cor       string `json:"cor"`
		UsuarioId uint   `json:"usuario_id"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Criar novo veículo
	veiculo := schemas.Veiculos{
		Placa:     input.Placa,
		Modelo:    input.Modelo,
		Cor:       input.Cor,
		UsuarioId: input.UsuarioId,
	}

	if err := db.Create(&veiculo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar veículo no banco de dados"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Veículo criado com sucesso"})
}
