package constants

var ExitCommand = map[string]bool{
	"\x1b": true,
	"exit": true,
	"quit": true,
	"stop": true,
}

//user command

const (
	NewUser                 = "new user"
	NewMerchant             = "new merchant"
	NewTxn                  = "new txn" // new transaction
	UpdateMerchant          = "update merchant"
	PayBack                 = "payback"
	ReportDiscount          = "report discount"
	ReportDues              = "report dues"
	ReportUserAtCreditLimit = "report users-at-credit-limit"
	ReportTotalDuse         = "report total-dues"
)
