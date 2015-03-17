package facade

import "testing"

type testpair struct {
	value   float64
	balance float64
}

var withdrawalTests = []testpair{
	{500, 9500},
	{9400, 100},
	{500, 100},
}

var depositTests = []testpair{
	{500, 10500},
	{9400, 19900},
	{500, 20400},
}

func TestWithdrawCash(t *testing.T) {
	var accessingBank BankAccountFacade = BuildBankAccountFacade(12345678, 1234)

	for _, pair := range withdrawalTests {
		accessingBank.withdrawCash(pair.value)
		if accessingBank.Balance() != pair.balance {
			t.Error(
				"For", pair,
				"expected", pair.balance,
				"got", accessingBank.Balance(),
			)
		}
	}
}

func TestDepositCash(t *testing.T) {
	var accessingBank BankAccountFacade = BuildBankAccountFacade(12345678, 1234)

	for _, pair := range depositTests {
		accessingBank.depositCash(pair.value)
		if accessingBank.Balance() != pair.balance {
			t.Error(
				"For", pair,
				"expected", pair.balance,
				"got", accessingBank.Balance(),
			)
		}
	}
}
