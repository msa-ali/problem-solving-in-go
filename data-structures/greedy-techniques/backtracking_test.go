package greedytechniques

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpGame(t *testing.T) {
	assert.Equal(t, true, jumpGame([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, false, jumpGame([]int{3, 2, 1, 0, 4}))
}

func TestRescueBoats(t *testing.T) {
	assert.Equal(t, 2, rescueBoats([]int{2, 1, 1}, 3))
	assert.Equal(t, 2, rescueBoats([]int{2, 1, 1}, 3))
}

func TestGasStations(t *testing.T) {
	assert.Equal(t, 3, gasStationJourney([]int{1, 2, 3, 4, 5}, []int{3, 4, 5, 1, 2}))
}

func TestTwoCityScheduling(t *testing.T) {
	assert.Equal(t, 350, twoCityScheduling([][]int{{100, 10}, {1000, 10}, {500, 50}, {100, 1}, {50, 1}, {2000, 40}}))
}
