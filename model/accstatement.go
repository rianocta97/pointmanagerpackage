package model

type StarStatementReq struct {
	AccountNumber string `json:"accountNumber"`
	StartDate     string `json:"startDate"`
	EndDate       string `json:"endDate"`
	Mode          string `json:"mode"`
}

type PartnerHistoryReq struct {
	AccountToken string `json:"accToken"`
}

type StarStatementResp struct {
	TransactionReferenceNumber string `json:"transactionReferenceNumber"`
	StatementId                string `json:"statementId"`
	EndingBalance              string `json:"endingBalance"`
	Description                string `json:"description"`
	PostDate                   string `json:"postDate"`
	DebitAmount                string `json:"debitAmount"`
	TrnsactionCode             string `json:"transactionCode"`
	Detail                     string `json:"detail"`
	CreditAmount               string `json:"creditAmount"`
	TransactionDate            string `json:"transactionDate"`
}

type PartnerHistoryResp struct {
	TransactionList []TransactionList `json:"transactionList"`
}

type TransactionList struct {
	Amount          float64 `json:"amount"`
	Partner         string  `json:"partner"`
	TrxId           string  `json:"trxId"`
	CustomerNote    string  `json:"customerNote"`
	Currency        string  `json:"currency"`
	TransactionDate string  `json:"transactionDate"`
}