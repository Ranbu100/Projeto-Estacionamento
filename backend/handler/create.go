package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/Ranbu100/Projeto-Estacionamento/utils"
	"github.com/gin-gonic/gin"
)

func CreateEntradasSaidasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando entrada/saidas de veiculos",
	})
}
func CreatePagamentosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando pagamentos",
	})
}
func CreateReservasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando reservas de vagas",
	})
}
func CreateUsuariosHandler(ctx *gin.Context) {

	var input struct {
		Nome  string `json:"nome"`
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	//JSON recebido
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando usuario",
	})
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}
	hashedPassword, err := utils.HashPassword(input.Senha)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao gerar hash da senha"})
		return
	}
	//salvando no bd a senha hasheada
	usuario := schemas.Usuarios{
		Nome:  input.Nome,
		Email: input.Email,
		Senha: hashedPassword,
	}
	if err := db.Create(&usuario).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"erro": "Erro ao salvar usuario no banco de dados"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Usuario criado com sucesso"})
}

func LoginHandler(ctx *gin.Context) {
	var input struct {
		Email string `json:"email"`
		Senha string `json:"senha"`
	}

	//JSON recebido
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}

	//buscar user por email
	var usuario schemas.Usuarios
	if err := db.Where("email = ?", input.Email).First(&usuario).Error; err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credencias invalidas"})
		return
	}

	//comparar senhas
	if !utils.CheckPasswordHash(input.Senha, usuario.Senha) {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Credencias invalidas"})
		return
	}

	//se a senha tiver correta
	ctx.JSON(http.StatusOK, gin.H{"message": "Login bem-sucedido"})
}

func CreateVagasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando vagas de estacionamento",
	})
}

func CreateVeiculosHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Criando veiculos",
	})
}
