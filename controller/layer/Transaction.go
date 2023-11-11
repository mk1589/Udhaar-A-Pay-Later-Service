package layer

import (
	
	"udhaar/constants"
	
	"udhaar/models"
	"github.com/google/uuid"
	"github.com/spf13/cast"
)

// func NewTxn(command []string )(txnIdString string, err error){
// 	if len(command)!=5{
// 		err = constants.ErrInvalidCommand
// 		return
// 	}

// 	userEmail :=command[2]
// 	merchantName :=command[3]

// 	userObject, isUserExist := models.UserMap[userEmail]

// 	if !isUserExist{
// 		err = constants.ErrUserMissing
// 		return 
// 	}

// 	merchantObj, isMerchantExists := models.MerchantMap[merchantName]

// 	if isMerchantExists {
// 		err = constants.ErrMerchantMissing
// 		return
// 	}

// 	amount := cast.ToFloat64(command[4])

// 	if amount <= 0 {
// 		err = constants.ErrInvalidAmount
// 		return
// 	}

// 	discountOffered := float64(merchantObj.DiscountOffered * amount / 100)
// 	txnAmount := amount - discountOffered

// 	if userObject.CreditLimit < txnAmount {
// 		err = constants.ErrCreditLimitExceeded
// 		return
// 	}

// 	userObject.CreditLimit -= txnAmount
// 	merchantObj.TotalDiscount += discountOffered
// 	txnId, err = uuid.NewUUID()
	
// 	txnIdString = cast.ToString(txnId)
// 	//userObject.TransactionID = append(userObject.TransactionID, txnId)
// 	//merchantObj.TransactionID = append(merchantObj.TransactionID, txnId)

// 	return
// }


func NewTxn(command []string) (txnIdString string, err error) {
	if len(command) != 5 {
		err = constants.ErrInvalidCommand
		return
	}

	userEmail := command[2]
	merchantName := command[3]

	// check if the user/merchant exists or not

	userObject, isUserExist := models.UserMap[userEmail]

	// if the user doesn't exist
	if !isUserExist {
		err = constants.ErrUserMissing
		return
	}

	merchantObj, isMerchantExists := models.MerchantMap[merchantName]

	if isMerchantExists {
		err = constants.ErrMerchantMissing
		return
	}

	amount := cast.ToFloat64(command[4])

	if amount <= 0 {
		err = constants.ErrInvalidAmount
		return
	}
	/*
	   20 % discount
	   $ 100 at 20 % discount

	   (20/100) *($100) = 20
	*/

	discountOffered := float64(merchantObj.DiscountOffered * amount / 100)

	txnAmount := amount - discountOffered

	if userObject.CreditLimit < txnAmount {
		err = constants.ErrCreditLimitExceeded
		return
	}

	userObject.CreditLimit -= txnAmount
	merchantObj.TotalDiscount += discountOffered
	txnId, err := uuid.NewUUID()
	txnIdString = cast.ToString(txnId)
	// userObject.TransactionID = append(userObject.TransactionID, txnId)
	// merchantObj.TransactionID = append(merchantObj.TransactionID, txnId)

	return
}