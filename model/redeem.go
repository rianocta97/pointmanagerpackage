package model

type StarRedeemReq struct {
	FromAccount    string `json:"fromAccount"`
	TrxAmt         string `json:"trxAmt"`
	TrxId          string `json:"trxId"`
	ItemName       string `json:"itemName"`
	Qty            string `json:"qty"`
	TrxType        string `json:"trxType"`
	BillerCode     string `json:"billerCode"`
	RequestDate    string `json:"requestDate"`
	CashbackAmount string `json:"cashbackAmount"`
}

type PartnerRedeemReq struct {
	FromAccToken string  `json:"fromAccToken"`
	FromMemberId string  `json:"fromMemberId"`
	FromSmmaId   string  `json:"fromSmmaId"`
	TrxId        string  `json:"trxId"`
	TrxAmt       float64 `json:"trxAmt"`
	ItemName     string  `json:"itemName"`
	Qty          float64 `json:"qty"`
	TrxType      string  `json:"trxType"`
	BillerCode   string  `json:"billerCode"`
	RequestDate  string  `json:"requestDate"`
}

type StarRedeemResp struct {
	TrxId       string  `json:"trxId"`
	TrxAmt      float64 `json:"trxAmt"`
	RequestDate string  `json:"requestDate"`
	FromAccount string  `json:"fromAccount"`
	BillerCode  string  `json:"billerCode"`
}

type PartnerRedeemResp struct {
	FromAccToken string  `json:"fromAccToken"`
	RequestDate  string  `json:"requestDate"`
	TrxAmt       float64 `json:"trxAmt"`
	TrxId        string  `json:"trxId"`
	BillerCode   string  `json:"billerCode"`
}