package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/Ranbu100/Projeto-Estacionamento/models"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Printf("Erro ao fazer decode do json: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := models.Insert(todo)

	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":    true,
			"Mensagem": fmt.Sprintf("ocorreu um erro ao tentar inserir: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":    false,
			"Mensagem": fmt.Sprintf("todo inserido com sucesso! ID: %d", id),
		}
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
