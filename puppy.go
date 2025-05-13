package puppy

import (
	"github.com/aditya020402/dog1"
)

func Bark() string {
	return "Woof!!"
}

func Barks() string {
	return "Woof !! Woof !! Woof !!"
}

func BigBark() string {
	return dog1.WhenGrownUp(Bark())
}

func BigBarks() string {
	return dog1.WhenGrownUp(Barks())
}

func Eating() string {
	return "Puppy is eating food."
}
