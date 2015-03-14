package strategy

import "testing"

func TestBird(t *testing.T) {
	bird := buildBird()
	exp := "Flying High"
	res := bird.tryToFly()

	if res != exp {
		t.Error(
			"For tryToFly()",
			"expected ", exp,
			"got", res,
		)
	}
}

func TestDog(t *testing.T) {
	dog := buildDog()
	exp := "I can't fly :("
	res := dog.tryToFly()

	if res != exp {
		t.Error(
			"For tryToFly()",
			"expected ", exp,
			"got", res,
		)
	}
}

func TestDogWithWings(t *testing.T) {
	dog := buildDog()
	dog.setFlyingAbility(CanFly{})
	exp := "Flying High"
	res := dog.tryToFly()

	if res != exp {
		t.Error(
			"For tryToFly()",
			"expected ", exp,
			"got", res,
		)
	}
}
