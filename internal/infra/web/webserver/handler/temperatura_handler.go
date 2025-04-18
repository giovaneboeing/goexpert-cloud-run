package handler

import (
	"encoding/json"
	"github.com/giovaneboeing/desafio-cloud-run/internal/infra/usecase"
	"net/http"
	"strings"
	"unicode"
)

type TemperaturaHandler struct {
	temperaturaUsecase *usecase.ConsultaTemperaturaUseCase
}

func NewTemperaturaHandler(temperaturaUsecase *usecase.ConsultaTemperaturaUseCase) *TemperaturaHandler {
	return &TemperaturaHandler{
		temperaturaUsecase: temperaturaUsecase,
	}
}

func (h *TemperaturaHandler) Consultar(w http.ResponseWriter, r *http.Request) {
	cep := extrairNumeros(r.URL.Query().Get("cep"))
	if len(cep) != 8 {
		http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
		return
	}

	outputDto, err := h.temperaturaUsecase.Execute(cep)
	if err != nil {
		if strings.Contains(err.Error(), "invalid zipcode") {
			http.Error(w, err.Error(), http.StatusUnprocessableEntity)
			return
		}
		if strings.Contains(err.Error(), "can not find zipcode") {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(outputDto)
}

func extrairNumeros(s string) string {
	var r []rune
	for _, c := range s {
		if unicode.IsDigit(c) {
			r = append(r, c)
		}
	}
	return string(r)
}
