package config

import ( 

	"os"

	"github.com/Ranbu100/Projeto-Estacionamento/schemas"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializePostgres() (*gorm.DB, error) {
	logger := GetLogger("postgres")
	dbPath := "./db/main.db" // Correção: remover tipo explícito na inicialização

	// Verificando se o DB existe
	_, err := os.Stat(dbPath) // Correção: remover `_ fs.FileInfo`
	if os.IsNotExist(err) {
		logger.Info("database não encontrado, criando...")

		// Criando o diretório
		err = os.MkdirAll("./db", os.ModePerm) // Correção: não usar `path:`, passe o valor diretamente
		if err != nil {
			return nil, err
		}

		// Criando o arquivo do banco de dados
		file, err := os.Create(dbPath) // Correção: não usar `nmae:`, passe o valor diretamente
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	// Conectando ao banco de dados Postgres
	dsn := "user=nicholas password=qwe123 dbname=estacionamento port=5432 sslmode=disable TimeZone=Brazil"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Errorf("postgres opening error: %v", err) // Corrigido para passar os parâmetros corretamente
		return nil, err
	}

	// Automigrate
	err = db.AutoMigrate(&schemas.Opening{}) // Passagem correta da estrutura para automigrate
	if err != nil {
		logger.Errorf("postgres automigration error: %v", err)
		return nil, err
	}

	// Retornando o banco de dados
	return db, nil
}
