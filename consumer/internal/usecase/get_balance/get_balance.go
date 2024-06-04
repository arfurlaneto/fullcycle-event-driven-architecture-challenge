package get_balance

import (
	"github.com.br/arfurlaneto/fullcycle-event-driven-architecture-challenge/consumer/internal/gateway"
)

type GetBalanceInputDTO struct {
	AccountId string  `json:"client_id"`
	Balance   float64 `json:"balance"`
}

type GetBalanceUseCase struct {
	BalanceGateway gateway.BalanceGateway
}

func NewGetBalanceUseCase(a gateway.BalanceGateway) *GetBalanceUseCase {
	return &GetBalanceUseCase{
		BalanceGateway: a,
	}
}

func (uc *GetBalanceUseCase) Execute(input GetBalanceInputDTO) (float64, error) {
	return uc.BalanceGateway.GetBalance(input.AccountId)
}
