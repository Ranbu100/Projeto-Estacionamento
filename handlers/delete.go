package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/Ranbu100/Projeto-Estacionamento/models"
	"github.com/go-chi/chi"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Erro ao fazer parse do id: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := models.Delete(int64(id))
	if err != nil {
		log.Printf("Erro ao deletar registro: %v", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	if rows > 1 {
		log.Printf("Error: foram deletados %d registros", rows)

	}

	resp := map[string]any{
		"Error":    false,
		"Mensagem": "dados deletados com sucesso",
	}

	w.Header().Add("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
