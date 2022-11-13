package constant

import customError "test/pkg/error"

var ApiKeyMissingError = customError.New("403", "", "ERR_API_KEY_MISSING", "API key is missing")
var InvalidApiKey = customError.New("401", "", "ERR_INVALID_API_KEY", "Invalid API key")
var InvalidEmail = customError.New("400", "", "ERR_INVALID_EMAIL", "Invalid Email")
var InvalidUserID = customError.New("400", "", "ERR_INVALID_USER_ID", "Invalid UserID")
var InvalidCreditCardData = customError.New("400", "", "ERR_INVALID_CREDIT_CARD_DATA", "Credit card data invalid")
var InvalidCreditCardFormat = customError.New("400", "", "ERR_INVALID_CREDIT_CARD_FORMAT", "Please provide *** fields.")
var InvalidCreditCardExpireFormat = customError.New("400", "", "ERR_INVALID_CREDIT_CARD_EXPIRE_FORMAT", "credit_card_expired format should be 02/2016")
var InvalidQueryOffsetAndLimitTypeData = customError.New("400", "", "ERR_INVALID_QUERY_OFFSET_TYPE_DATA", `invalid query param "of" and "lt" type data. "of" and "lt"" must be number`)
var UserNotFound = customError.New("400", "", "ERR_USER_NOT_FOUND", "User not found ")
var SystemError = customError.New("500", "", "ERR_SYSTEM_ERROR", "Something went wrong. Please try again later.")

func CustomBadRequest(err error) error {
	return customError.New("400", "", "ERR_BAD_REQUEST", err.Error())
}
