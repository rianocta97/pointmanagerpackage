package model

type GetBalanceParam struct {
	AccountNumber string `json:"accountNumber"`
}

type InquiryBalanceParam struct {
	AccountToken string `json:"accToken"`
}