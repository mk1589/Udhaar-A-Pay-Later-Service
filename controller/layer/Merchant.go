package layer

import (
	"fmt"
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
