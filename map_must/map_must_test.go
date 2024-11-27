package map_must_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/map_must"
)

func TestEquals(t *testing.T) {
	map_must.Equals(map[int]string{
		1: "a",
		2: "b",
	}, map[int]string{
		2: "b",
		1: "a",
	})

	tests.ExpectPanic(t, func() {
		map_must.Equals(map[string]int{
			"a": 1,
		}, map[string]int{
			"b": 2,
		})
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, map_must.Nice(map[string]int{"c": 3, "b": 2, "a": 1}))

	tests.ExpectPanic(t, func() {
		map_must.Nice(map[string]int{})
	})
}
