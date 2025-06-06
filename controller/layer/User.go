package layer

import (
	"fmt"
	"udhaar/constants"
	"udhaar/models"
	"udhaar/utils"

	"github.com/spf13/cast"
)

func AddUser(command []string) (newUserName string, err error) {

	if len(command) != 5 {
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
	creditLimit, err := utils.IsValidCreditLimit(command[4])
	_, ok := models.UserMap[email]
	if ok {
		err = constants.ErrUserAlreadyExists
		return
	}

	newUser := models.User{
		Name:        name,
		Email:       email,
		CreditLimit: creditLimit,
	}

	models.UserMap[name] = newUser

	newUserName = newUser.Name
	return
}

func GetUserAtCreditLimit(command []string) (listAsString string, err error) {

	if len(command) != 2 {
		err = constants.ErrInvalidCommand
		return
	}

	if len(models.UserMap) > 0 {
		for _, user := range models.UserMap {
			if user.Dues == user.CreditLimit {
				listAsString += fmt.Sprintf("\n%v %v", user.Name, user.Email)
			}
		}

	} else {
		err = constants.ErrUserListEmpty
	}

	if len(listAsString) > 0 {
		listAsString = fmt.Sprintf(
			"The list of users at credit limit is as follows: \n%v",
			listAsString,
		)
	} else {
		listAsString = "There is no user at limit right now\n"
	}
	return
}

func ShowUserDues(command []string) (duesString string, err error) {
	/*
		report dues <user-email>
	*/

	if len(command) != 3 {
		err = constants.ErrInvalidCommand
		return
	}

	// check if the user/merchant exists or not

	userName := command[2]
	userObject, isUserExist := models.UserMap[userName]

	// if the user doesn't exist
	if !isUserExist {
		err = constants.ErrUserMissing
		return
	}
	duesString = cast.ToString(userObject.Dues)

	return
}

func GetTotalDues(command []string) (totalDuesString string, err error) {
	totalDues := 0.0
	// sum	the total dues of the users

	for _, userObj := range models.UserMap {
		totalDues += userObj.Dues
	}

	totalDuesString = cast.ToString(totalDues)
	return
}
