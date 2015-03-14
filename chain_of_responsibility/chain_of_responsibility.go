package changeofresponsibility

import "fmt"
import "errors"

type Numbers struct {
	Number1           float64
	Number2           float64
	CalculationWanted string
}

func BuildNumbers(number1 float64, number2 float64, calc string) Numbers {
	return Numbers{number1, number2, calc}
}

type Chain interface {
	setNextChain(Chain)
	calculate(Numbers) (float64, error)
}

type Chainable struct {
	nextInChain Chain
}

func (c *Chainable) setNextChain(nextChain Chain) {
	c.nextInChain = nextChain
}

type AddNumbers struct {
	Chainable
}

func (a AddNumbers) calculate(request Numbers) (float64, error) {
	if request.CalculationWanted == "add" {
		result := (request.Number1 + request.Number2)
		fmt.Println(request.Number1, " + ", request.Number2, " = ", result)

		return result, nil
	}

	if a.nextInChain == nil {
		return 0, errors.New("Out of chain")
	} else {
		return a.nextInChain.calculate(request)
	}
}

type SubtractNumbers struct {
	Chainable
}

func (s SubtractNumbers) calculate(request Numbers) (float64, error) {
	if request.CalculationWanted == "sub" {
		result := (request.Number1 - request.Number2)
		fmt.Println(request.Number1, " - ", request.Number2, " = ", result)
		return result, nil
	}

	if s.nextInChain == nil {
		return 0, errors.New("Out of chain")
	} else {
		return s.nextInChain.calculate(request)
	}
}

type MultiplyNumbers struct {
	Chainable
}

func (m MultiplyNumbers) calculate(request Numbers) (float64, error) {
	if request.CalculationWanted == "mul" {
		result := (request.Number1 * request.Number2)
		fmt.Println(request.Number1, " * ", request.Number2, " = ", result)
		return result, nil
	}

	if m.nextInChain == nil {
		return 0, errors.New("Out of chain")
	} else {
		return m.nextInChain.calculate(request)
	}
}

type DivideNumbers struct {
	Chainable
}

func (d DivideNumbers) calculate(request Numbers) (float64, error) {
	if request.CalculationWanted == "div" {
		result := (request.Number1 / request.Number2)
		fmt.Println(request.Number1, " / ", request.Number2, " = ", result)
		return result, nil
	}

	if d.nextInChain == nil {
		return 0, errors.New("Out of Chain")
	} else {
		return d.nextInChain.calculate(request)
	}
}
