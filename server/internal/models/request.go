package models

type AuthRequest struct {
	Email    string `json:"login"`
	Password string `json:"pass"`
}

type RegRequest struct {
	Email    string `json:"login"`
	Password string `json:"pass"`
}

type DepositeRequest struct {
	Amount string `json:"amount"`
}

type TransferRequest struct {
	Login  string `json:"login"`
	Amount string `json:"amount"`
}
