package map_must

import (
	"testing"

	"github.com/yyle88/must/internal/utils"
)

func TestEquals(t *testing.T) {
	Equals(map[int]string{
		1: "a",
		2: "b",
	}, map[int]string{
		2: "b",
		1: "a",
	})

	utils.ExpectPanic(t, func() {
		Equals(map[string]int{
			"a": 1,
		}, map[string]int{
			"b": 2,
		})
	})
}
