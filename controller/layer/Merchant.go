package layer

import (
	"fmt"
	"github.com/spf13/cast"
	"udhaar/utils"
	"udhaar/constants"
	"udhaar/models"
)

func AddMerchant(command []string) (newMerchantDetails string, err error) {
	if len(command) != 4 {
		err = constants.ErrInvalidCommand
		return
	}
	name, err := utils.IsValidName(command[2])
	if err != nil {
		return
	}
	discount, err := utils.IsValidDiscount(command[3])
	if err != nil {
		return
	}

	_, ok := models.MerchantMap[name]
	if ok {
		err = constants.ErrMerchantAlreadyExists
		return
	}

	newMerchant := models.Merchant{
		UniqueName:      name,
		DiscountOffered: discount,
	}

	models.MerchantMap[name] = newMerchant

	newMerchantDetails = fmt.Sprintf(
		"%v (%v)",
		newMerchant.UniqueName,
		newMerchant.DiscountOffered,
	)
	return
}

func UpdateMerchant(command []string) (newDisountString string, err error) {

	/*
		update merchant m1 1%
	*/

	if len(command) != 4 {
		err = constants.ErrInvalidCommand
		return
	}

	name := command[2]


	merchantObj, isExists := models.MerchantMap[name]
	if !isExists {
		err = constants.ErrMerchantMissing
		return
	}

	newDiscountFloat, err := utils.IsValidDiscount(command[3])
	if err != nil {
		return
	}

	merchantObj.DiscountOffered = newDiscountFloat
	models.MerchantMap[name] = merchantObj
	newDisountString = cast.ToString(newDiscountFloat)
	return
}

func GetDiscount(command []string) (discountString string, err error) {

	/*
		report discount m1
	*/
	if len(command) != 3 {
		err = constants.ErrInvalidCommand
		return
	}

	merchantName := command[2]


	merchantObject, isUserExist := models.MerchantMap[merchantName]


	if !isUserExist {
		err = constants.ErrMerchantMissing
		return
	}
	discountString = cast.ToString(merchantObject.DiscountOffered)

	return

}