package handler

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Ranbu100/Projeto-Estacionamento/enums"
	"github.com/Ranbu100/Projeto-Estacionamento/middleware"
	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/Ranbu100/Projeto-Estacionamento/utils"
	"github.com/Ranbu100/Projeto-Estacionamento/validator"
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

	//Validando Entrada e saida
	entradaSaidaData := validator.EntradaSaida{
		VeiculoId:   input.VeiculoID,
		DataEntrada: input.Entrada,
		DataSaida:   *input.Saida,
	}

	if err := validator.ValidateEntradaSaida(entradaSaidaData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		UsuarioId        uint    `json:"Usuario_Id"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}

	//Validando pagamento
	pagamentoData := validator.Pagamento{
		Valor:     input.Valor,
		UsuarioId: input.UsuarioId,
	}

	if err := validator.ValidatePagamento(pagamentoData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		VagaId      uint      `json:"vaga_id"`
		UsuarioId   uint      `json:"usuario_id"`
		DataHora    time.Time `gorm:"not null" json:"data_hora_reserva"`
		Duracao     uint      `gorm:"not null" json:"duracao"`
		Status      string    `gorm:"size:50;not null" json:"status_reserva"`
		DataReserva string    `json:"DataReserva"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	//Validando reservas
	reservationData := validator.Reserva{
		UsuarioId:   input.UsuarioId,
		VagaId:      input.VagaId,
		DataReserva: input.DataReserva,
	}

	if err := validator.ValidateReserva(reservationData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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

	userData := validator.User{
		Nome:     input.Nome,
		Email:    input.Email,
		Telefone: input.Telefone,
	}
	//Validar a struct de entrada (input)
	if err := validator.ValidateUser(userData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
		log.Println("Erro no JSON recebido:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verificar se as credenciais são de um admin
	adminEmail := os.Getenv("ADMIN_EMAIL")
	adminSenha := os.Getenv("ADMIN_SENHA")

	// Caso seja admin, checar as credenciais diretamente
	if input.Email == adminEmail && input.Senha == adminSenha {
		token, err := middleware.GenerateJWT("admin", "admin") // Adiciona o papel
		if err != nil {
			log.Println("Erro ao gerar token:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
			return
		}
		log.Println("Login de admin bem-sucedido")
		ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "Login de admin bem-sucedido!"})
		return
	}

	// Caso contrário, verificar no banco de dados para usuários comuns
	var usuario schemas.Usuarios
	if err := db.Where("email = ?", input.Email).First(&usuario).Error; err != nil {
		log.Println("Usuário não encontrado:", input.Email)
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Comparar senhas
	if !utils.CheckPasswordHash(input.Senha, usuario.Senha) {
		log.Println("Falha na comparação da senha")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credenciais inválidas"})
		return
	}

	// Gerar token JWT para o usuário comum com a role
	token, err := middleware.GenerateJWT(usuario.Email, usuario.Role) // Inclui a role do usuário
	if err != nil {
		log.Println("Erro ao gerar token para o usuário:", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar token"})
		return
	}

	// Se a senha estiver correta
	log.Println("Login bem-sucedido para usuário:", usuario.Email)
	ctx.JSON(http.StatusOK, gin.H{"token": token, "message": "Login bem-sucedido"})
}

func CreateVagasHandler(c *gin.Context) {
	var input struct {
		NumeroVaga int   `json:"numero_vaga"`
		TipoVaga   uint8 `json:"tipo_vaga"` // Altere o tipo para uint8 para corresponder à enumeração
		Status     int   `json:"status_vaga"`
	}

	// Verificar se o JSON foi recebido corretamente
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Verificar se já existe uma vaga com o mesmo número
	var vagaExistente schemas.Vagas
	if err := db.Where("numero_vaga = ?", input.NumeroVaga).First(&vagaExistente).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Já existe uma vaga com este número"})
		return
	}

	// Validação do tipo de vaga usando os enums
	if !enums.TipoVagaValido(input.TipoVaga) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de vaga inválido"})
		return
	}

	// Validação do status da vaga usando os enums
	if !enums.StatusVagaValido(input.Status) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Status de vaga inválido"})
		return
	}

	// Mapear o tipo de vaga: enums.TipoCarro = 1 (carro), enums.TipoMoto = 2 (moto)
	var tipoVaga uint8
	switch input.TipoVaga {
	case 1:
		tipoVaga = enums.TipoCarro
	case 2:
		tipoVaga = enums.TipoMoto
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "Tipo de vaga inválido"})
		return
	}

	// Criar nova vaga
	vaga := schemas.Vagas{
		NumeroVaga: input.NumeroVaga,
		TipoVaga:   tipoVaga,
		Status:     input.Status,
	}

	// Salvar no banco de dados
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

	// Verificar se o JSON foi recebido corretamente
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos"})
		return
	}

	// Validação dos dados do veículo
	veiculoData := validator.Veiculo{
		Modelo: input.Modelo,
		Placa:  input.Placa,
	}

	// Validar o veículo usando o pacote de validação
	if err := validator.ValidateVeiculo(veiculoData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Criar novo veículo com os dados fornecidos
	veiculo := schemas.Veiculos{
		Placa:     input.Placa,
		Modelo:    input.Modelo,
		Cor:       input.Cor,
		UsuarioId: input.UsuarioId,
	}

	// Salvar veículo no banco de dados
	if err := db.Create(&veiculo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar veículo no banco de dados"})
		return
	}

	// Retornar sucesso
	c.JSON(http.StatusCreated, gin.H{"message": "Veículo criado com sucesso"})
}
