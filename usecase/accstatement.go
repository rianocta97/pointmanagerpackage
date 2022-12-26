package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rianocta97/pointmanagerpackage/constants"
	"github.com/rianocta97/pointmanagerpackage/model"
)

type AccStatementUsecase interface {
	ValidateRequestData(dataStar *model.StarStatementReq, dataPartner *model.PartnerHistoryReq) error
	PostData(dataStar *model.StarStatementReq, dataPartner *model.PartnerHistoryReq, coreApi string) (*[]model.StarStatementResp, *model.PartnerHistoryResp, error)
}

type AccStatementUsecaseImpl struct{}

func (as *AccStatementUsecaseImpl) ValidateRequestData(dataStar *model.StarStatementReq, dataPartner *model.PartnerHistoryReq) error{
	if dataStar != nil {
		if dataStar.AccountNumber == "" || len(dataStar.AccountNumber) > 50 {
			return errors.New(constants.ErrorValidate_InvalidAccountNumber)
		}
		if dataStar.Mode == "" || len(dataStar.Mode) > 5 {
			return errors.New(constants.ErrorValidate_InvalidMode)
		}
		return nil
	} else if dataPartner != nil {
		if dataPartner.AccountToken == "" {
			return errors.New(constants.ErrorValidate_InvalidAccountToken)
		}
		return nil
	}
	return errors.New(constants.Error_EmptyReqBody)
}

func (as *AccStatementUsecaseImpl) PostData(dataStar *model.StarStatementReq, dataPartner *model.PartnerHistoryReq, coreApi string) (*[]model.StarStatementResp, *model.PartnerHistoryResp, error){
	var postBody []byte
	if dataStar != nil {
		postBody, _ = json.Marshal(dataStar)
	} else if dataPartner != nil {
		postBody, _ = json.Marshal(dataPartner)
	} else {
		return nil,nil, errors.New(constants.Error_EmptyReqBody)
	}

	reqBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(coreApi, "application/json", reqBody)
	if err != nil {
		return nil, nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, nil, err
	}

	if dataStar != nil {
		var dataStar *[]model.StarStatementResp
		if err := json.Unmarshal(body, &dataStar); err != nil {
			return nil, nil, err
		}
		return dataStar, nil, nil
	} else if dataPartner != nil {
		dataPartner := &model.PartnerHistoryResp{}
		if err := json.Unmarshal(body, dataPartner); err != nil {
			return nil, nil, err
		}
		return nil, dataPartner, nil
	}
	return nil, nil, errors.New(constants.Error_SendRequest)
}

func InitAccStatementUsecase() AccStatementUsecase {
	return &AccStatementUsecaseImpl{}
}