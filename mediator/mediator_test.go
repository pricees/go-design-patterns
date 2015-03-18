package mediator

import (
	"fmt"
	"testing"
)

type testpair struct {
	values  []float64
	average float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

func TestAverage(t *testing.T) {
	nyse := BuildStockMediator()
	broker := BuildGormanSlacks(nyse)
	broker2 := BuildJTPoorman(nyse)

	broker.saleOffer("MSFT", 100)
	broker.saleOffer("GOOG", 50)

	broker2.buyOffer("MSFT", 100)
	broker2.saleOffer("NRG", 10)

	broker.buyOffer("NRG", 10)

	fmt.Println("================")
	nyse.getStockOfferings()

	/*
		for _, pair := range tests {
			v := Average(pair.values)
			if v != pair.average {
				t.Error(
					"For", pair.values,
					"expected", pair.average,
					"got", v,
				)
			}
		}*/
}
