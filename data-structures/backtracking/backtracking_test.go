package backtracking

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpGame(t *testing.T) {
	assert.Equal(t, true, jumpGame([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, false, jumpGame([]int{3, 2, 1, 0, 4}))
}
