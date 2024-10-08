package validator

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Nome     string `validate:"required,min=3,max=50"`
	Email    string `validate:"required,email"`
	Telefone string `validate:"required,len=11"`
}

type Veiculo struct {
	Modelo string `validate:"required,min=2,max=50"`
	Placa  string `validate:"required,len=7"`
}

type Vagas struct {
	Tipo   uint8 `validade:"required,oneof=1 2"`
	Status int   `validate:"required,oneof=1 2 3 4"`
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

	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Nome":
				return fmt.Errorf("nome deve ter entre 3 e 50 caracteres")
			case "Email":
				return fmt.Errorf("email deve ser valido")
			case "Telefone":
				return fmt.Errorf("telefone deve ter 9 digitos e conter apenas numeros")
			}
		}
	}
	return err
}

func ValidateVeiculo(veiculo Veiculo) error {
	validate := validator.New()

	// Registrar validação customizada para a placa
	validate.RegisterValidation("placa_valid", PlacaValidator)

	err := validate.Struct(veiculo)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Modelo":
				return fmt.Errorf("modelo deve ter entre 2 e 50 caracteres")
			case "Placa":
				return fmt.Errorf("placa inválida, deve ter 7 caracteres no formato correto")
			}
		}
	}
	return err
}

// Validação customizada para placas de veículos
func PlacaValidator(fl validator.FieldLevel) bool {
	placa := fl.Field().String()
	// Expressão regular para validar formato de placa
	re := regexp.MustCompile(`^[A-Z]{3}\d{4}$`)
	return re.MatchString(placa)
}

func StatusValidator(f1 validator.FieldLevel) bool {
	status := f1.Field().Int()
	validStatuses := []int{1, 2, 3, 4}

	for _, v := range validStatuses {
		if status == int64(v) {
			return true
		}

	}
	return false
}

func TipoVagaValidator(f1 validator.FieldLevel) bool {
	tipo := f1.Field().Int()
	validTipos := []int{1, 2} // 1: carro, 2: moto

	for _, v := range validTipos {
		if tipo == int64(v) {
			return true
		}
	}
	return false
}

func ValidateVaga(vaga Vagas) error {
	validate := validator.New()

	// Registrar validações personalizadas
	validate.RegisterValidation("status_valid", StatusValidator)
	validate.RegisterValidation("tipo_valid", TipoVagaValidator)

	if err := validate.Struct(vaga); err != nil {
		return err
	}
	return nil
}

func ValidateEntradaSaida(entradaSaida EntradaSaida) error {
	validate := validator.New()

	err := validate.Struct(entradaSaida)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "DataEntrada":
				return fmt.Errorf("data de entrada deve estar no formato YYYY-MM-DD HH:MM:SS")
			case "DataSaida":
				return fmt.Errorf("data de saída deve ser maior que a data de entrada e estar no formato correto")
			}
		}
	}
	return err
}

func ValidatePagamento(pagamento Pagamento) error {
	validate := validator.New()

	err := validate.Struct(pagamento)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Valor":
				return fmt.Errorf("valor deve ser maior ou igual a zero")
			}
		}
	}
	return err
}

func ValidateReserva(reserva Reserva) error {
	validate := validator.New()

	err := validate.Struct(reserva)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "DataReserva":
				return fmt.Errorf("data da reserva deve estar no formato YYYY-MM-DD HH:MM:SS")
			}
		}
	}
	return err
}
