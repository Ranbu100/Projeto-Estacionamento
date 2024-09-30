package validator

import (
	"github.com/go-playground/validator/v10"
)

type User struct {
	Nome     string `validate:"required,min=3,max=50"`
	Email    string `validate:"required,email"`
	Telefone string `validate:"required,len-9"`
}

type Veiculo struct {
	Modelo string `validate:"required,min=2,max=50"`
	Placa  string `validate:"required,len=7"`
}

type Vagas struct {
	Tipo   string `validade:"required,oneof=normal especial"`
	Status string `validate:"required"`
}

type EntradaSaida struct {
	VeiculoId   uint   `validate:"required"`
	DataEntrada string `validate:"required"`
	DataSaida   string `validate:"required"`
}

type Pagamento struct {
	UsuarioId uint    `validate:"required"`
	Valor     float64 `validate:"required,min-0"`
}

type Reserva struct {
	UsuarioId   uint   `validate:"required"`
	VagaId      uint   `validate:"required"`
	DataReserva string `validate:"required"`
}

func ValidateUser(user User) error {
	validate := validator.New()
	return validate.Struct(user)
}

func ValidateVeiculo(veiculo Veiculo) error {
	validate := validator.New()
	return validate.Struct(veiculo)
}

func ValidateVaga(vaga Vagas) error {
	validate := validator.New()
	return validate.Struct(vaga)
}

func ValidateEntradaSaida(entradaSaida EntradaSaida) error {
	validate := validator.New()
	return validate.Struct(entradaSaida)
}

func ValidatePagamento(pagamento Pagamento) error {
	validate := validator.New()
	return validate.Struct(pagamento)
}

func ValidateReserva(reserva Reserva) error {
	validate := validator.New()
	return validate.Struct(reserva)
}
