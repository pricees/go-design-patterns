package strategy

import "fmt"

type Flys interface {
	fly() string
}

type CanFly struct{}

func (i CanFly) fly() string {
	return "Flying High"
}

type CantFly struct{}

func (c CantFly) fly() string {
	return "I can't fly :("
}

type Animal struct {
	Name    string
	Height  float64
	weight  int
	favFood string
	speed   float64
	sound   string

	flyingType Flys
}

func (a Animal) tryToFly() string {
	return a.flyingType.fly()
}

func (a *Animal) setFlyingAbility(newFlyType Flys) {
	a.flyingType = newFlyType
}

type Dog struct {
	Animal
}

func (d Dog) digHole() {
	fmt.Println("Dug a hole")
}

func buildDog() Dog {
	return Dog{Animal{sound: "Bark", flyingType: CantFly{}}}
}

type Bird struct {
	Animal
}

func buildBird() Bird {
	return Bird{Animal{sound: "Tweet", flyingType: CanFly{}}}
}
