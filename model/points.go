package model

type StarGetBalanceRequestParam struct {
	AccountNumber string `json:"accountNumber"`
}

type PartnerGetBalanceRequestParam struct {
	AccountToken string `json:"accToken"`
}

type StarResponseData struct {
	AccountNumber string  `json:"accountNumber"`
	Balance       float64 `json:"balance"`
	Name          string  `json:"name"`
	StarData      bool    `json:"starData"`
}

type PartnerResponseData struct {
	AccountNumber string  `json:"accountNumber"`
	Balance       float64 `json:"balance"`
	Name          string  `json:"name"`
}