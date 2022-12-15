package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rianocta97/pointmanagerpackage/constant"
	"github.com/rianocta97/pointmanagerpackage/model"
)

type PointUsecase interface {
	CheckGetBalanceData(rawData interface{}) (*model.StarGetBalanceRequestParam, *model.PartnerGetBalanceRequestParam, error)
	validateRequestData(dataStar *model.StarGetBalanceRequestParam, dataPartner *model.PartnerGetBalanceRequestParam) error
	PostData(dataStar *model.StarGetBalanceRequestParam, dataPartner *model.PartnerGetBalanceRequestParam) (map[string]interface{}, error)
	CheckGetBalanceResponse(rawData map[string]interface{}) (*model.StarResponseData, *model.PartnerResponseData, error)
	validateResponseData(dataStar *model.StarResponseData, dataPartner *model.PartnerResponseData) error
}

type PointUsecaseImpl struct {
}

func (p *PointUsecaseImpl) CheckGetBalanceData(rawData interface{}) (*model.StarGetBalanceRequestParam, *model.PartnerGetBalanceRequestParam, error) {
	data := rawData.(map[string]interface{})
	if _, ok := data["accountNumber"]; ok {
		dataStar := &model.StarGetBalanceRequestParam{
			AccountNumber: data["accountNumber"].(string),
		}
		if err := p.validateRequestData(dataStar, nil); err != nil {
			return nil, nil, err
		}
		return dataStar, nil, nil
	} else if _, ok := data["accToken"]; ok {
		dataPartner := &model.PartnerGetBalanceRequestParam{
			AccountToken: data["accToken"].(string),
		}
		if err := p.validateRequestData(nil, dataPartner); err != nil {
			return nil, nil, err
		}
		return nil, dataPartner, nil
	} 

	return nil, nil, errors.New(constant.ErrorValidate_BalanceData)
}

func (p *PointUsecaseImpl) validateRequestData(dataStar *model.StarGetBalanceRequestParam, dataPartner *model.PartnerGetBalanceRequestParam) error{
	// assumption: length will be 10 max
	if dataStar != nil {
		if (len(dataStar.AccountNumber) <= 0) || (len(dataStar.AccountNumber) > 10) {
			return errors.New(constant.ErrorValidate_InvalidAccountNumber)
		}
		return nil
	} else if dataPartner != nil {
		if (len(dataPartner.AccountToken) <= 0) || (len(dataPartner.AccountToken) > 10) {
			return errors.New(constant.ErrorValidate_InvalidAccountToken)
		}
		return nil
	}
	return errors.New(constant.ErrorValidate_BalanceData)
}

func (p *PointUsecaseImpl) PostData(dataStar *model.StarGetBalanceRequestParam, dataPartner *model.PartnerGetBalanceRequestParam) (map[string]interface{}, error) {
	var postBody []byte
	if dataStar != nil {
		postBody, _ = json.Marshal(dataStar)
	} else if dataPartner != nil {
		postBody, _ = json.Marshal(dataPartner)
	} else {
		return nil, errors.New(constant.ErrorValidate_BalanceData)
	}

	reqBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:3000/core/getbalance", "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responses := make(map[string]interface{})
	json.Unmarshal(body, &responses)
	return responses, nil
}

func (p *PointUsecaseImpl) CheckGetBalanceResponse(rawData map[string]interface{}) (*model.StarResponseData, *model.PartnerResponseData, error) {
	if _, ok := rawData["starData"]; ok {
		dataStar := &model.StarResponseData{
			AccountNumber: rawData["accountNumber"].(string),
			Balance: rawData["balance"].(float64),
			Name: rawData["name"].(string),
			StarData: rawData["starData"].(bool),
		}
		if err := p.validateResponseData(dataStar, nil); err != nil {
			return nil, nil, err
		}
		return dataStar, nil, nil
	} else {
		// in this case, there's no difference between star & partner response data.
		// so the "starData" param is added to differentiate between them
		dataPartner := &model.PartnerResponseData{
			AccountNumber: rawData["accountNumber"].(string),
			Balance: rawData["balance"].(float64),
			Name: rawData["name"].(string),
		}
		if err := p.validateResponseData(nil, dataPartner); err != nil {
			return nil, nil, err
		}
		return nil, dataPartner, nil
	}
}

func (p *PointUsecaseImpl) validateResponseData(dataStar *model.StarResponseData, dataPartner *model.PartnerResponseData) error{
	// assumption: length will be 20 max
	if dataStar != nil {
		if (len(dataStar.AccountNumber) <= 0) || (len(dataStar.AccountNumber) > 20) {
			return errors.New(constant.ErrorValidate_InvalidAccountNumber)
		}
		return nil
	} else if dataPartner != nil {
		if (len(dataPartner.AccountNumber) <= 0) || (len(dataPartner.AccountNumber) > 20) {
			return errors.New(constant.ErrorValidate_InvalidAccountNumber)
		}
		return nil
	}
	return errors.New(constant.ErrorValidate_BalanceData)
}

func InitPointUsecase() PointUsecase {
	return &PointUsecaseImpl{}
}