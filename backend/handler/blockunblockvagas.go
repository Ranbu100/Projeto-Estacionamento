package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/gin-gonic/gin"
)

func BlockUnblockVagaHandler(c *gin.Context) {
	id := c.Param("id")
	var vaga schemas.Vagas

	if err := db.First(&vaga, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vaga não encontrada"})
		return
	}

	var input struct {
		Status string `json:"status"` // Status pode ser: Bloqueada, disponivel, reservada etc
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados invalidos"})
		return
	}

	//Atualizar o status da vaga
	vaga.Status = input.Status
	if err := db.Save(&vaga).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar vaga"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaga atualizada com sucesso"})
}
