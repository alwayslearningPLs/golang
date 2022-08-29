package directcomposition

type Athlete struct{}

func (a Athlete) Train() string {
	return "trainning"
}

type Swimmer struct {
	Athlete Athlete
	Swim    func() string
}

func Swim() string {
	return "swimming"
}

func NewSwimmer() Swimmer {
	return Swimmer{Swim: Swim} // zero-initialization, which means that Athlete is a struct, so we don't need to initialize manually.
}

type Animal struct{}

func (a Animal) Eat() string {
	return "eating"
}

type Fish struct {
	Animal Animal
	Swim   func() string
}

func NewFish() Fish {
	return Fish{Swim: Swim} // zero-initialization, which means that Animal is a struct, so we don't need to initialize manually.
}
