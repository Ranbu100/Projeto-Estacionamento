package config

import (
	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")

	// Conectando ao banco de dados Postgres no AntaresSQL
	dsn := "host=localhost user=postgres password=qwer1234 dbname=bd_estacionamento port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres opening error: %v", err)
		return nil, err
	}

	// Automigrate para as structs corretas
	err = db.AutoMigrate(&schemas.Usuarios{}, &schemas.Vagas{}, &schemas.Veiculos{}, &schemas.EntradasSaidas{}, &schemas.Pagamentos{}, &schemas.Tarifas{})
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}

	// Retornando o banco de dados
	return db, nil
}
