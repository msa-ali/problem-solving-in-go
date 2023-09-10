package fp

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

type Item struct {
	id    int
	value int
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	itemMapper := Map(func(item int, index int) *Item {
		return &Item{
			id:    index,
			value: item * 2,
		}
	})
	result := itemMapper(arr)
	expected := []*Item{
		{
			id:    0,
			value: 2,
		},
		{
			id:    1,
			value: 4,
		},
		{
			id:    2,
			value: 6,
		},
	}
	require.True(t, reflect.DeepEqual(result, expected))
}
