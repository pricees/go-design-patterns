package facade

import "fmt"

type WelcomeToBank struct{}

func BuildWelcomeToBank() WelcomeToBank {
	fmt.Println("Welcome to ABC Bank")
	return WelcomeToBank{}
}

//////////

type AccountNumberCheck struct {
	accountNumber int
}

func BuildAccountNumberCheck(accountNumber int) AccountNumberCheck {
	if accountNumber == 0 {
		accountNumber = 12345678 // default
	}
	return AccountNumberCheck{accountNumber}
}

func (anc AccountNumberCheck) getAccountNumber() int {
	return anc.accountNumber
}

func (anc AccountNumberCheck) accountActive(acctNumToCheck int) bool {
	return anc.getAccountNumber() == acctNumToCheck
}

//////////

type SecurityCodeCheck struct {
	securityCode int
}

func BuildSecurityCodeCheck(securityCode int) SecurityCodeCheck {
	if securityCode == 0 {
		securityCode = 1234 // default
	}
	return SecurityCodeCheck{securityCode}
}

func (scc SecurityCodeCheck) getSecurityCode() int {
	return scc.securityCode
}

func (scc SecurityCodeCheck) isCodeCorrect(secCodeToCheck int) bool {
	return scc.getSecurityCode() == secCodeToCheck
}

//////////

type FundsCheck struct {
	cashInAccount float64
}

func (fc *FundsCheck) getCashInAccount() float64 {
	return fc.cashInAccount
}

func (fc *FundsCheck) decreaseCashInAccount(cashWithdrawn float64) {
	fc.cashInAccount -= cashWithdrawn
}

func (fc *FundsCheck) increaseCashInAccount(cashDeposited float64) {
	fc.cashInAccount += cashDeposited
}

func (fc *FundsCheck) hasEnoughMoney(cashToWithdrawal float64) bool {
	if cashToWithdrawal > fc.getCashInAccount() {
		fmt.Println("Error: No dinero")
		fmt.Println("Current Balance: ", fc.getCashInAccount())
		return false
	} else {
		fc.decreaseCashInAccount(cashToWithdrawal)
		fmt.Println("Withdrawal completed")
		fmt.Println("Current Balance: ", fc.getCashInAccount())
		return true
	}
}

func (fc *FundsCheck) makeDeposit(cashToDeposit float64) {
	fc.increaseCashInAccount(cashToDeposit)
	fmt.Println("Deposit completed")
	fmt.Println("Current Balance: ", fc.getCashInAccount())
}

func BuildFundsCheck(cashInAccount float64) *FundsCheck {
	return &FundsCheck{cashInAccount}
}

//////////

type BankAccountFacade struct {
	accountNumber int
	securityCode  int

	acctCheck   AccountNumberCheck
	codeCheck   SecurityCodeCheck
	fundCheck   *FundsCheck
	bankWelcome WelcomeToBank
}

func BuildBankAccountFacade(accountNumber int, securityCode int) BankAccountFacade {
	return BankAccountFacade{
		accountNumber,
		securityCode,
		BuildAccountNumberCheck(accountNumber),
		BuildSecurityCodeCheck(securityCode),
		BuildFundsCheck(10000),
		BuildWelcomeToBank(),
	}
}

func (baf BankAccountFacade) getAccountNumber() int {
	return baf.accountNumber
}

func (baf BankAccountFacade) getSecurityCode() int {
	return baf.securityCode
}

func (baf BankAccountFacade) withdrawCash(cashToGet float64) {
	if baf.authorized() && baf.fundCheck.hasEnoughMoney(cashToGet) {
		fmt.Println("Transaction Complete.")
	} else {
		fmt.Println("Transaction Failed.")
	}
}

func (baf BankAccountFacade) depositCash(cashToDeposit float64) {
	if baf.authorized() {
		baf.fundCheck.makeDeposit(cashToDeposit)
		fmt.Println("Transaction Complete.")
	} else {
		fmt.Println("Transaction Failed.")
	}
}

func (baf BankAccountFacade) Balance() float64 {
	return baf.fundCheck.getCashInAccount()
}
func (baf BankAccountFacade) authorized() bool {
	return baf.acctCheck.accountActive(baf.getAccountNumber()) &&
		baf.codeCheck.isCodeCorrect(baf.getSecurityCode())
}
