package handler

import (
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"github.com/gin-gonic/gin"
)

func UnblockVagasHandler(c *gin.Context) {
	id := c.Param("id")
	var vaga schemas.Vagas

	if err := db.First(&vaga, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Vaga nao encontrada"})
		return
	}

	if vaga.Status != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Vaga nao esta bloqueada"})
		return
	}

	vaga.Status = 1
	if err := db.Save(&vaga).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao desbloquear vaga"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Vaga desbloqueada com suecesso"})
}
