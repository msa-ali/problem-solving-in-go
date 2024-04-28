package stacks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator(t *testing.T) {
	result := Calculator("4+(15-12)+100")
	assert.Equal(t, 107, result)
}
