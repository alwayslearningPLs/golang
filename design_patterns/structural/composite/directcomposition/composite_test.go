package directcomposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwimmer(t *testing.T) {
	var (
		wantTrain = "trainning"
		wantSwim  = "swimming"
	)

	got := NewSwimmer()

	assert.Equal(t, wantTrain, got.Athlete.Train())
	assert.Equal(t, wantSwim, got.Swim())
}

func TestFish(t *testing.T) {
	var (
		wantEat  = "eating"
		wantSwim = "swimming"
	)

	got := NewFish()

	assert.Equal(t, wantEat, got.Animal.Eat())
	assert.Equal(t, wantSwim, got.Swim())
}
