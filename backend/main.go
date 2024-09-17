package main

import (
	"github.com/Ranbu100/Projeto-Estacionamento/config"
	"github.com/Ranbu100/Projeto-Estacionamento/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//inciando o config
	err := config.Init()
	if err != nil {
		logger.Errorf("config initialization err: %v", err) // Passando a string de formato e o valor diretamente
		return
	}

	//Inicializando o roteamento
	router.Initialize()
}
