package interfacecomposition

// Athlete ... I don't know which verb to use here as someone that trains...
type Athlete interface {
	Train() string
}

type AthleteImpl struct{}

func (a AthleteImpl) Train() string {
	return "trainning"
}

type Swimmer interface {
	Swim() string
}

type SwimmerImpl struct{}

func (s SwimmerImpl) Swim() string {
	return "swimming"
}

type SwimmerAthlete struct {
	Athlete
	Swimmer
}

func NewSwimmerAthlete() SwimmerAthlete {
	return SwimmerAthlete{Athlete: AthleteImpl{}, Swimmer: SwimmerImpl{}}
}

type Eater interface {
	Eat() string
}

type Animal struct{}

func (a Animal) Eat() string {
	return "eating"
}

type Fish struct {
	Eater
	Swimmer
}

func NewFish() Fish {
	return Fish{Eater: Animal{}, Swimmer: SwimmerImpl{}}
}
