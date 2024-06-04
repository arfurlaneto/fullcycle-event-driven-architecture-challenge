package gateway

type BalanceGateway interface {
	GetBalance(clientID string) (float64, error)
	UpdateBalance(clientID string, balance float64) error
}
