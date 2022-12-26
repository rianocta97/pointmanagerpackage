package constants

const (
	Error_EmptyReqBody  = "request body not found"
	Error_EmptyRespBody = "request body not found"

	ErrorValidate_InvalidAccountNumber = "invalid account number"
	ErrorValidate_InvalidAccountToken  = "invalid account token"
	ErrorValidate_InvalidTrxAmount     = "invalid transaction amount"
	ErrorValidate_InvalidTrxId         = "invalid transaction id"
	ErrorValidate_InvalidBillerCode    = "invalid biller code"
	ErrorValidate_InvalidRequestDate   = "invalid request date"
	ErrorValidate_InvalidMode          = "invalid mode"

	Error_SendRequest = "error sending request"
)

func IsNotANumberError(param string) string {
	return "error " + param + " is not a number"
}