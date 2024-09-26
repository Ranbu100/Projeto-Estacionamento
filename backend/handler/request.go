package handler

import (
	"github.com/Ranbu100/Projeto-Estacionamento/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetPostgres()

	logger.Info("Hanlder incializado")
}
