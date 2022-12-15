package pointmanagerpackage

import (
	"encoding/json"

	"github.com/rianocta97/pointmanagerpackage/usecase"
)

func CheckGetBalanceData(data interface{}) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if err := json.Unmarshal(data.([]byte), &m); err != nil {
		return nil, err
	}

	pu := usecase.InitPointUsecase()
	dataStar, dataPartner, err := pu.CheckGetBalanceData(m)
	if err != nil {
		return nil, err
	}
	responses, err := pu.PostData(dataStar, dataPartner)
	if err != nil {
		return nil, err
	}

	print(responses, err)

	return responses, nil
}