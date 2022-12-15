package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rianocta97/pointmanagerpackage/constant"
	"github.com/rianocta97/pointmanagerpackage/model"
)

type PointUsecase interface {
	CheckGetBalanceData(rawData interface{}) (*model.GetBalanceParam, *model.InquiryBalanceParam, error)
	validateGetBalanceData(dataStar *model.GetBalanceParam, dataPartner *model.InquiryBalanceParam) error
	PostData(dataStar *model.GetBalanceParam, dataPartner *model.InquiryBalanceParam) (map[string]interface{}, error)
}

type PointUsecaseImpl struct {
}

func (p *PointUsecaseImpl) CheckGetBalanceData(rawData interface{}) (*model.GetBalanceParam, *model.InquiryBalanceParam, error) {
	data := rawData.(map[string]interface{})
	if _, ok := data["accountNumber"]; ok {
		dataStar := &model.GetBalanceParam{
			AccountNumber: data["accountNumber"].(string),
		}
		fmt.Println("datastar: ", dataStar)
		if err := p.validateGetBalanceData(dataStar, nil); err != nil {
			return nil, nil, err
		}
		return dataStar, nil, nil
	} else if _, ok := data["accToken"]; ok {
		dataPartner := &model.InquiryBalanceParam{
			AccountToken: data["accToken"].(string),
		}
		fmt.Println("datapartner: ", dataPartner)
		if err := p.validateGetBalanceData(nil, dataPartner); err != nil {
			return nil, nil, err
		}
		return nil, dataPartner, nil
	} 

	return nil, nil, errors.New(constant.ErrorValidate_BalanceData)
}

func (p *PointUsecaseImpl) validateGetBalanceData(dataStar *model.GetBalanceParam, dataPartner *model.InquiryBalanceParam) error{
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

func (p *PointUsecaseImpl) PostData(dataStar *model.GetBalanceParam, dataPartner *model.InquiryBalanceParam) (map[string]interface{}, error) {
	var postBody []byte
	if dataStar != nil {
		postBody, _ = json.Marshal(dataStar)
	} else if dataPartner != nil {
		postBody, _ = json.Marshal(dataStar)
	} else {
		return nil, errors.New(constant.ErrorValidate_BalanceData)
	}

	reqBody := bytes.NewBuffer(postBody)
	resp, err := http.Post("http://localhost:3000/core/getbalance", "application/json", reqBody)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(reqBody)
	if err != nil {
		return nil, err
	}

	responses := make(map[string]interface{})
	json.Unmarshal(body, &responses)
	fmt.Println(responses)
	return responses, nil
}

func InitPointUsecase() PointUsecase {
	return &PointUsecaseImpl{}
}