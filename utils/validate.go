package utils

import(
	"strings"
	"udhaar/constants"
	"github.com/spf13/cast"
)

 func IsValidEmail(email string) (string, error ){
	email=strings.ReplaceAll(email, " ", "")
	if len(email) == 0{
		return "", constants.ErrEmptyEmail
	}
	return email, nil
 }

 func IsValidName(name string) (string, error) {
	name=strings.ReplaceAll(name, " ", "")
	if len(name) == 0{
		return "", constants.ErrEmptyName
	}
	return name, nil
 }

 func IsValidCreditLimit(creditLimitString string) (float64,error) {
	creditLimitString = strings.ReplaceAll(creditLimitString, " ", "")
	if len(creditLimitString) == 0{
		return 0.0, constants.ErrEmptyCreditLimit
	}

	creditLimitFloat :=cast.ToFloat64(creditLimitString)
	if creditLimitFloat <=0.0{
		return 0.0, constants.ErrEmptyCreditLimit
	}
	return creditLimitFloat, nil
 }

 func IsValidDiscount(discountString string) (float64, error) {
	if discountString[len(discountString)-1] != '%' {
		return 0.0, constants.ErrInvalidDiscount
	}
	discount := cast.ToFloat64(discountString)
	if discount <= 0 || discount > 100 {
		return 0.0, constants.ErrInvalidDiscount
	}
	return discount, nil
}