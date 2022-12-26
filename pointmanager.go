package pointmanagerpackage

import (
	"github.com/rianocta97/pointmanagerpackage/model"
	"github.com/rianocta97/pointmanagerpackage/usecase"
)

func CheckGetBalanceData(dataStar *model.StarGetBalanceReq, dataPartner *model.PartnerGetBalanceReq, endPoint string) (*model.GetBalanceResp, error) {
	pu := usecase.InitGetBalanceUsecase()
	if err := pu.ValidateRequestData(dataStar, dataPartner); err != nil {
		return nil, err
	}

	responses, err := pu.PostData(dataStar, dataPartner, endPoint)
	if err != nil {
		return nil, err
	}

	response, err := pu.CheckResponseData(responses)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func AccStatement(dataStar *model.StarStatementReq, dataPartner *model.PartnerHistoryReq, endPoint string) (*[]model.StarStatementResp, *model.PartnerHistoryResp,error){
	as := usecase.InitAccStatementUsecase()
	if err := as.ValidateRequestData(dataStar, dataPartner); err != nil {
		return nil,nil,err
	}

	responseStar, responsePartner, err := as.PostData(dataStar, dataPartner, endPoint)
	if err != nil {
		return nil,nil,err
	}

	return responseStar, responsePartner, nil
}

func GenericRedeem(dataStar *model.StarRedeemReq, dataPartner *model.PartnerRedeemReq, endPoint string) (*model.StarRedeemResp, *model.PartnerRedeemResp, error) { 
	ru := usecase.InitRedeemUsecase()

	if err := ru.ValidateRequestData(dataStar, dataPartner); err != nil {
		return nil, nil, err
	}

	responses, err := ru.PostData(dataStar, dataPartner, endPoint)
	if err != nil {
		return nil, nil, err
	}

	responseStar, responsePartner, err := ru.CheckResponse(responses)
	if err != nil {
		return nil, nil, err
	}

	return responseStar, responsePartner, nil
}