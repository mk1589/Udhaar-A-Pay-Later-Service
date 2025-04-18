package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"udhaar/constants"
	"udhaar/controller/layer"
	"udhaar/utils"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println(
		constants.UdhaarAppSymbol + " \nwelcome to the Udhaar - Buy Now Pay Later Application \nPlease enter your input: ")
	for {
		utils.PromptMsg()
		input, err := reader.ReadString('\n')

		input = strings.Replace(input, "\n", "", -1)
		input = strings.Trim(input, " ")

		if err != nil {
			fmt.Print(err)
			os.Exit(0)
		}

		if _, ok := constants.ExitCommand[input]; ok {
			fmt.Println(constants.Red + "Program Terminated" + constants.White)
			fmt.Println(
				constants.Reset + "Thank you for Exploring Pay-Later Terminal Client" + constants.White,
			)
			os.Exit(0)
		}

		command := strings.Split(input, " ")
		var rootCommand string

		if command[0] != "payback" {
			rootCommand = strings.Join(command[:2], " ")
		} else {
			rootCommand = command[0]
		}

		var arg string
		// var err error
		var operationCode int
		switch rootCommand {
		case constants.NewUser:
			arg, err = layer.AddUser(command)
			operationCode = constants.AddUser
		case constants.NewMerchant:
			arg, err = layer.AddMerchant(command)
			operationCode = constants.AddMerchant
		case constants.NewTxn:
			arg, err = layer.NewTxn(command)
			operationCode = constants.Txn
		case constants.UpdateMerchant:
			arg, err = layer.UpdateMerchant(command)
			operationCode = constants.UpdateDiscount
		case constants.PayBack:
			arg, err = layer.PayBack(command)
			operationCode = constants.Payback
		case constants.ReportDiscount:
			arg, err = layer.GetDiscount(command)
			operationCode = constants.GetTotalMerchantDiscount
		case constants.ReportDues:
			arg, err = layer.ShowUserDues(command)
			operationCode = constants.GetUserDues
		case constants.ReportUserAtCreditLimit:
			arg, err = layer.GetUserAtCreditLimit(command)
			operationCode = constants.GetUsersAtLimit

		case constants.ReportTotalDuse:
			arg, err = layer.GetTotalDues(command)
			operationCode = constants.GetTotalDues
		default:
			err = constants.ErrInvalidCommand
		}

		utils.PrintMsg(err, operationCode, arg)
	}

	//fmt.Println("the input provider is ", input)
}
