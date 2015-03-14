package changeofresponsibility

import "testing"

type testpair struct {
	values []float64
	calc   string
	result float64
}

var tests = []testpair{
	{[]float64{1, 2}, "add", 3},
	{[]float64{1, 2}, "sub", -1},
	{[]float64{1, 2}, "mul", 2},
	{[]float64{1, 2}, "div", 0.5},
	{[]float64{1, 2}, "foo", 0},
}

func TestChangeOfResponsibility(t *testing.T) {
	var addCalc Chain = &AddNumbers{}
	var subCalc Chain = &SubtractNumbers{}
	var mulCalc Chain = &MultiplyNumbers{}
	var divCalc Chain = &DivideNumbers{}

	addCalc.setNextChain(subCalc)
	subCalc.setNextChain(mulCalc)
	mulCalc.setNextChain(divCalc)

	for _, pair := range tests {

		request := Numbers{pair.values[0], pair.values[1], pair.calc}
		result, _ := addCalc.calculate(request)
		if result != pair.result {
			t.Error(
				"For", pair.values,
				"expected", pair.result,
				"got", result,
			)
		}
	}
}
