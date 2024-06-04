package web

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/usecase/get_balance"
	"github.com/go-chi/chi/v5"
)

type WebBalanceHandler struct {
	GetBalanceUseCase get_balance.GetBalanceUseCase
}

func NewWebBalanceHandler(getBalanceUseCase get_balance.GetBalanceUseCase) *WebBalanceHandler {
	return &WebBalanceHandler{
		GetBalanceUseCase: getBalanceUseCase,
	}
}

func (h *WebBalanceHandler) GetBalance(w http.ResponseWriter, r *http.Request) {
	accountId := chi.URLParam(r, "account_id")
	var dto = get_balance.GetBalanceInputDTO{
		AccountId: accountId,
	}

	output, err := h.GetBalanceUseCase.Execute(dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
