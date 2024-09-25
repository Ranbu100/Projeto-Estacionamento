package schemas

import "time"

type Usuarios struct {
	Id          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Nome        string `gorm:"size:100;not null" json:"nome"`
	Email       string `gorm:"size:100;not null;unique" json:"email"`
	Senha       string `gorm:"size:100;not null" json:"senha"`
	TipoUsuario string `gorm:"size:50;not null" json:"tipo_usuario"`
	Telefone    string `gorm:"size:50;not null" json:"telefone"`
}

type Vagas struct {
	ID         uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	NumeroVaga int    `gorm:"not null" json:"numero_vaga"`
	TipoVaga   string `gorm:"size:50;not null" json:"tipo_vaga"`
	Status     string `gorm:"size:50;not null" json:"status_vaga"`
}

type Veiculos struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Placa     string `gorm:"size:10;not null;unique" json:"placa"`
	Modelo    string `gorm:"size:100" json:"modelo"`
	Cor       string `gorm:"size:40" json:"cor"`
	UsuarioId uint   `gorm:"not null" json:"usuario_id"`
}

type EntradasSaidas struct {
	ID        uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	VeiculoId uint       `gorm:"not null" json:"veiculo_id"`
	VagaId    uint       `gorm:"not null" json:"vaga_id"`
	Entrada   time.Time  `gorm:"not null" json:"horario_entrada"`
	Saida     *time.Time `json:"horario_saida,omitempty"` // Campo opcional, pois a saída pode ser nula
}

type Pagamentos struct {
	ID               uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	EntradasSaidasId uint    `gorm:"not null" json:"entradas_saidas_id"`
	Valor            float64 `gorm:"type:decimal(10,2);not null" json:"valor_pago"`
	Metodo           string  `gorm:"size:50;not null" json:"metodo_pagamento"`
	Status           string  `gorm:"size:50;not null" json:"status_pagamento"`
}

type Tarifas struct {
	ID        uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Descricao string  `gorm:"size:100;not null" json:"descricao"`
	Valor     float64 `gorm:"type:decimal(10,2);not null" json:"valor"`
}
type Reservas struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UsuarioId uint      `gorm:"not null" json:"usuario_id"`
	VagaId    uint      `gorm:"not null" json:"vaga_id"`
	DataHora  time.Time `gorm:"not null" json:"data_hora_reserva"`
	Duracao   uint      `gorm:"not null" json:"duracao"`
	Status    string    `gorm:"size:50;not null" json:"status_reserva"`
}
