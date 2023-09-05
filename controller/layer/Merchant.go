package layer

import (
	"udhaar/utils"
	"udhaar/constants"

)

func AddUser(command []string) (newUserName string, err error){

	if len(command) != 5{
		err = constants.ErrInvalidCommand
		return
	}
	name, err := utils.IsValidName(command[2])

	if err != nil {
		return
	}
	email, err := utils.IsValidEmail(command[3])
	if err != nil {
		return
	}
	creditLimit,err := utils.IsValidCreditLimit(command[4])
	_,ok := models.UserMap[email]
	if ok{
		err=constants.ErrUserAlreadyExists
		return
	}

	newUser := models.User{
		Name: name,
		Email: email,
		CreditLimit: creditLimit,
	}

	models.UserMap[email]=newUser

	newUserName = newUser.Name
	return
}