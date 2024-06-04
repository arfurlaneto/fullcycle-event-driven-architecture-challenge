package update_balance

import (
	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/gateway"
)

type UpdateBalanceInputDTO struct {
	AccountId string  `json:"client_id"`
	Balance   float64 `json:"balance"`
}

type UpdateBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewUpdateBalanceUseCase(a gateway.BalanceGateway) *UpdateBalanceUseCase {
	return &UpdateBalanceUseCase{
		BalanceGateway: a,
	}
}

func (uc *UpdateBalanceUseCase) Execute(input UpdateBalanceInputDTO) error {
	return uc.BalanceGateway.UpdateBalance(input.AccountId, input.Balance)
}
