package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/Ranbu100/Projeto-Estacionamento/utils"
	"github.com/gin-gonic/gin"
)

func UpdateEntradasSaidasHandler(ctx *gin.Context) {

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
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando reservas",
	})
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
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando veiculo",
	})
}
func UpdateVagasHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "Atualizando vagas de estacionamento",
	})
}
