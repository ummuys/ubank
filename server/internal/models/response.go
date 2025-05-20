package models

type AuthResponse struct {
	Message string `json:"message"`
	Token   string `json:"token"`
}

type RegResponse struct {
	Message string `json:"message"`
}

type DepositeResponse struct {
	Message string `json:"message"`
}

type BalanceResponse struct {
	Message string `json:"message"`
}

type TransferResponse struct {
	Message string `json:"message"`
}
