package slice_must_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/slice_must"
)

func TestEquals(t *testing.T) {
	slice_must.Equals([]int{1, 2, 3}, []int{1, 2, 3})

	tests.ExpectPanic(t, func() {
		slice_must.Equals([]string{"a"}, []string{"b"})
	})
}

func TestContains(t *testing.T) {
	slice_must.Contains([]string{"a", "b", "c"}, "a")

	tests.ExpectPanic(t, func() {
		slice_must.Contains([]int{1, 2, 3}, 4)
	})
}

func TestIn(t *testing.T) {
	slice_must.In("a", []string{"a", "b", "c"})

	tests.ExpectPanic(t, func() {
		slice_must.In(4, []int{1, 2, 3})
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, []int{1, 2, 3}, slice_must.Nice([]int{1, 2, 3}))

	tests.ExpectPanic(t, func() {
		slice_must.Nice([]string{})
	})
}
