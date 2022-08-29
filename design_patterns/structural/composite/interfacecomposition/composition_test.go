package interfacecomposition

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSwimmerAthlete(t *testing.T) {
	var (
		wantTrain = "trainning"
		wantSwim  = "swimming"
	)

	got := NewSwimmerAthlete()

	assert.Equal(t, wantTrain, got.Train())
	assert.Equal(t, wantTrain, got.Athlete.Train())
	assert.Equal(t, wantSwim, got.Swim())
	assert.Equal(t, wantSwim, got.Swimmer.Swim())
}

func TestFish(t *testing.T) {
	var (
		wantEat  = "eating"
		wantSwim = "swimming"
	)

	got := NewFish()

	assert.Equal(t, wantEat, got.Eat())
	assert.Equal(t, wantEat, got.Eater.Eat())
	assert.Equal(t, wantSwim, got.Swim())
	assert.Equal(t, wantSwim, got.Swimmer.Swim())
}
