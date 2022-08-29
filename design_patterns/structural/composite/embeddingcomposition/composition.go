package embeddingcomposition

type Athlete struct{}

func (a Athlete) Train() string {
	return "trainning"
}

type Swimmer struct {
	Athlete // Embedded composition
	Swim    func() string
}

func Swim() string {
	return "swimming"
}

func NewSwimmer() Swimmer {
	return Swimmer{Swim: Swim}
}

type Animal struct{}

func (a Animal) Eat() string {
	return "eating"
}

type Fish struct {
	Animal // Embedded composition, so we can call directly to the functions/fields inside Animal struct
	Swim   func() string
}

func NewFish() Fish {
	return Fish{Swim: Swim}
}
