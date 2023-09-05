package constants

var ExitCommand = map[string]bool{
	"\x1b": true,
	"exit": true,
	"quit": true,
	"stop": true,
}


//user command


const(
	NewUser = "new user"
	NewMerchant="new merchant"
	NewTxn="new txn" // new transaction
	UpdateMerchant ="update merchant"
	PayBack ="payback"
	ReportDiscount="report dicount"
	ReportDues ="report dues"
	ReportUserAtCreditLimit="report user-at-credit-limit"
	ReportTotalDuse = "report total-dues"
)
