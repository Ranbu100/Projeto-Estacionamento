package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//Carregando as variaveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Fatal("Erro ao carregar arquivo .env")
	}

	//inicalizando postgres
	db, err = InitializePostgres()

	if err != nil {
		return fmt.Errorf("erro ao inicializar postgres: %v", err)
	}

	return nil
}

func GetPostgres() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	//Inicilizando o logger
	logger = NewLogger(p)
	return logger
}
