package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/rianocta97/pointmanagerpackage/constants"
	"github.com/rianocta97/pointmanagerpackage/model"
)

type RedeemUsecase interface {
	ValidateRequestData(dataStar *model.StarRedeemReq, dataPartner *model.PartnerRedeemReq) error
	PostData(dataStar *model.StarRedeemReq, dataPartner *model.PartnerRedeemReq, endPoint string) (map[string]interface{}, error)
	CheckResponse(rawData map[string]interface{}) (*model.StarRedeemResp, *model.PartnerRedeemResp, error)
}

type RedeemUsecaseImpl struct{}

func (r *RedeemUsecaseImpl) ValidateRequestData(dataStar *model.StarRedeemReq, dataPartner *model.PartnerRedeemReq) error{
	if dataStar != nil {
		if len(dataStar.FromAccount) > 45 {
			return errors.New(constants.ErrorValidate_InvalidAccountNumber)
		}
		if _, err := strconv.Atoi(dataStar.TrxAmt); err != nil {
			return errors.New(constants.IsNotANumberError("transaction amount"))
		}
		if len(dataStar.TrxId) > 16 {
			return errors.New(constants.ErrorValidate_InvalidTrxId)
		}
		if _, err := strconv.Atoi(dataStar.BillerCode); err != nil {
			return errors.New(constants.IsNotANumberError("biller code"))
		}
		if len(dataStar.BillerCode) > 10 {
			return errors.New(constants.ErrorValidate_InvalidBillerCode)
		}
		if _, err := time.Parse("02-01-2006 15:04:05", dataStar.RequestDate); err != nil {
    		return errors.New(constants.ErrorValidate_InvalidRequestDate)
		}
		return nil
	} else if dataPartner != nil {
		if dataPartner.TrxAmt <= 0 {
			return errors.New(constants.ErrorValidate_InvalidTrxAmount)
		}
		if len(dataPartner.TrxId) > 16 {
			return errors.New(constants.ErrorValidate_InvalidTrxId)
		}
		if len(dataPartner.BillerCode) > 10 {
			return errors.New(constants.ErrorValidate_InvalidBillerCode)
		}
		if _, err := time.Parse("02-01-2006 15:04:05", dataPartner.RequestDate); err != nil {
    		return errors.New(constants.ErrorValidate_InvalidRequestDate)
		}
		return nil
	}
	return errors.New(constants.Error_EmptyReqBody)
}

func (r *RedeemUsecaseImpl) PostData(dataStar *model.StarRedeemReq, dataPartner *model.PartnerRedeemReq, endPoint string) (map[string]interface{}, error) {
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	responses := make(map[string]interface{})
	json.Unmarshal(body, &responses)
	return responses, nil	
}

func (r *RedeemUsecaseImpl) CheckResponse(rawData map[string]interface{}) (*model.StarRedeemResp, *model.PartnerRedeemResp, error) {
	if _, ok := rawData["fromAccount"]; ok {
		jsonString, _ := json.Marshal(rawData)
		dataStar := &model.StarRedeemResp{}
		json.Unmarshal(jsonString, dataStar)
		return dataStar, nil, nil
	} else if _, ok := rawData["fromAccToken"]; ok {
		jsonString, _ := json.Marshal(rawData)
		dataPartner := &model.PartnerRedeemResp{}
		json.Unmarshal(jsonString, dataPartner)
		return nil, dataPartner, nil
	}
	return nil, nil, errors.New(constants.Error_EmptyRespBody)
}

func InitRedeemUsecase() RedeemUsecase {
	return &RedeemUsecaseImpl{}
}