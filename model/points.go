package model

type GetBalanceParam struct {
	AccountNumber string `json:"accountNumber"`
}

type InitBalanceParam struct {
	AccountToken string `json:"accToken"`
}