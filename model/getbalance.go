package model

type StarGetBalanceReq struct {
	AccountNumber string `json:"accountNumber"`
}

type PartnerGetBalanceReq struct {
	AccountToken string `json:"accToken"`
}

// since there's no difference between response, we'll use the same model for all
type GetBalanceResp struct {
	AccountNumber string  `json:"accountNumber"`
	Balance       float64 `json:"balance"`
	Name          string  `json:"name"`
}

// in case of change in the future, we could use the struct below and add/modify the property

// type StarGetBalanceResp struct {
// 	AccountNumber string  `json:"accountNumber"`
// 	Balance       float64 `json:"balance"`
// 	Name          string  `json:"name"`
// }

// type PartnerGetBalanceResp struct {
// 	AccountNumber string  `json:"accountNumber"`
// 	Balance       float64 `json:"balance"`
// 	Name          string  `json:"name"`
// }