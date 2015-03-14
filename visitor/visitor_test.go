package visitor

import "testing"

type testpair struct {
	goodForSale Visitable
	tax         float64
}

var tests = []testpair{
	{BuildLiquor(2.0), 0.36},
	{BuildTobacco(10), 3.2},
	{BuildNecessity(100.0), 0},
}

func TestTaxVisitor(t *testing.T) {
	var taxCalc TaxVisitor = TaxVisitor{}

	for _, pair := range tests {
		res := pair.goodForSale.Accept(taxCalc)

		if res != pair.tax {
			t.Error(
				"For TaxVisitor : ",
				"expected", res,
				"got", pair.tax,
			)
		}
	}
}
