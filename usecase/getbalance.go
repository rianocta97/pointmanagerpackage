package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/rianocta97/pointmanagerpackage/constants"
	"github.com/rianocta97/pointmanagerpackage/model"
)

type GetBalanceUsecase interface {
	ValidateRequestData(dataStar *model.StarGetBalanceReq, dataPartner *model.PartnerGetBalanceReq) error
	PostData(dataStar *model.StarGetBalanceReq, dataPartner *model.PartnerGetBalanceReq, endPoint string) (map[string]interface{}, error)
	CheckResponseData(rawData map[string]interface{}) (*model.GetBalanceResp, error)
}

type GetBalanceUsecaseImpl struct {}

func (gb *GetBalanceUsecaseImpl) ValidateRequestData(dataStar *model.StarGetBalanceReq, dataPartner *model.PartnerGetBalanceReq) error{
	if dataStar != nil {
		if (len(dataStar.AccountNumber) <= 0) || (len(dataStar.AccountNumber) > 50) {
			return errors.New(constants.ErrorValidate_InvalidAccountNumber)
		}
		return nil
	} else if dataPartner != nil {
		if (len(dataPartner.AccountToken) <= 0) || (len(dataPartner.AccountToken) > 100) {
			return errors.New(constants.ErrorValidate_InvalidAccountToken)
		}
		return nil
	}
	return errors.New(constants.Error_EmptyReqBody)
}

func (gb *GetBalanceUsecaseImpl) PostData(dataStar *model.StarGetBalanceReq, dataPartner *model.PartnerGetBalanceReq, endPoint string) (map[string]interface{}, error) {
	var postBody []byte
	if dataStar != nil {
		postBody, _ = json.Marshal(dataStar)
	} else if dataPartner != nil {
		postBody, _ = json.Marshal(dataPartner)
	} else {
		return nil, errors.New(constants.Error_EmptyReqBody)
	}

	reqBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(endPoint, "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// body, err := ioutil.ReadAll(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responses := make(map[string]interface{})
	json.Unmarshal(body, &responses)
	return responses, nil
}

func (gb *GetBalanceUsecaseImpl) CheckResponseData(rawData map[string]interface{}) (*model.GetBalanceResp, error) {
	// In this case, there's no difference between star & partner response data.
	// So for now, we're using the same model for response data.
	jsonString, _ := json.Marshal(rawData)
	data := &model.GetBalanceResp{}
	if err := json.Unmarshal(jsonString, data); err != nil {
		return nil, err
	}
	return data, nil

	// If there's any changes in the response data, we could use conditional if-else to differentiate the response
}

func InitGetBalanceUsecase() GetBalanceUsecase {
	return &GetBalanceUsecaseImpl{}
}