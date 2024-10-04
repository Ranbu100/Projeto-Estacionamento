package enums

const (
	TipoCarro = 1
	TipoMoto  = 2
)

const (
	Disponivel = 1
	Ocupada    = 2
	Reservada  = 3
	Manutencao = 4
)

func TipoVagaValido(tipo uint8) bool {
	return tipo == TipoCarro || tipo == TipoMoto
}

func StatusVagaValido(status int) bool {
	return status == Disponivel || status == Ocupada || status == Reservada || status == Manutencao
}
