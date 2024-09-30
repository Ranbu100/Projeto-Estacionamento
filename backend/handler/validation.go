package handler

import (
	"time"

	"github.com/go-playground/validator/v10"
)

// Estrutura de entrada para a criação de usuário
type CreateUserInput struct {
	Nome     string `json:"nome" validate:"required,min=2,max=100"`
	Email    string `json:"email" validate:"required,email"`
	Senha    string `json:"senha" validate:"required,min=6"`
	Telefone string `json:"telefone" validate:"required"`
}

type CreateReservasInput struct {
	VagaId      int       `json:"vagaid" validate:"required"`
	UsuarioId   int       `json:"usuarioid" validate:"required"`
	DataReserva time.Time `json:"data_reserva" validate:"required,datetime=2006-01-02 15:04:05"`
}

type CreateVagasInput struct {
	NumeroVaga int    `json:"numerovaga" validate:"required"`
	TipoVaga   string `json:"tipovaga" validate:"required"`
	Status     string `json:"status" validate:"required"`
}

type CreatePagamentosInput struct {
	EntradasSaidasId int     `json:"entrada-saidasid" validate:"required"`
	Valor            float64 `json:"valor" validate:"required"`
	Metodo           string  `json:"metodo" validate:"required"`
	Status           string  `json:"status" validate:"required"`
}

type CreateVeiculosInput struct {
	Placa     string `json:"placa" validate:"required"`
	Modelo    string `json:"modelo" validate:"required"`
	Cor       string `json:"cor" validate:"required"`
	UsuarioId int    `json:"usuarioid" validate:"required"`
}

// Instância do validador
var validate = validator.New()
