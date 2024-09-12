package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/models"
)

func List(w http.ResponseWriter, r *http.Request) {
	todos, err := models.GetAll()
	if err != nil {
		log.Printf("erro ao obter dados: %v", err)
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(todos)
}
