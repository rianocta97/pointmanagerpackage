package pointmanagerpackage

import (
	"encoding/json"

	"github.com/rianocta97/pointmanagerpackage/model"
	"github.com/rianocta97/pointmanagerpackage/usecase"
)

func CheckGetBalanceData(data interface{}) (*model.StarResponseData, *model.PartnerResponseData, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data.([]byte), &m); err != nil {
		return nil, nil, err
	}

	pu := usecase.InitPointUsecase()
	dataStar, dataPartner, err := pu.CheckGetBalanceData(m)
	if err != nil {
		return nil, nil, err
	}

	responses, err := pu.PostData(dataStar, dataPartner)
	if err != nil {
		return nil, nil, err
	}
	responseStar, responsePartner, err := pu.CheckGetBalanceResponse(responses)
	if err != nil {
		return nil, nil, err
	}

	return responseStar, responsePartner, err
}

// func main() {
// 	jsonData := `{"accountNumber": "00001"}`
// 	// jsonData := `{"accToken": "token01"}`

// 	bytes := []byte(jsonData)

// 	a, b, c:= CheckGetBalanceData(bytes)
// 	fmt.Println(a)
// 	fmt.Println(b)
// 	fmt.Println(c)
// }