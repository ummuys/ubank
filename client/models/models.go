package models

type User struct {
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type TransferRequest struct {
	Login  string `json:"login"`
	Amount string `json:"amount"`
}

type DepositeRequest struct {
	Amount string `json:"amount"`
}
