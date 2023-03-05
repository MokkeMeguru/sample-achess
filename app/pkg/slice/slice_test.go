package slice_test

import (
	"testing"

	"github.com/MokkeMeguru/simple-achess/app/pkg/slice"
	"github.com/stretchr/testify/require"
)

func Test_Map(t *testing.T) {
	type Custom struct {
		X int
		Y int
	}
	slice1 := []*Custom{
		{X: 10, Y: 10},
		{X: 10, Y: 15},
		{X: 8, Y: 10},
	}
	ret := slice.Map(slice1, func(s *Custom) int {
		return s.X * s.Y
	})
	require.Equal(t, []int{100, 150, 80}, ret)
}
