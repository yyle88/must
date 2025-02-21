package muststrings_test

import (
	"testing"

	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/muststrings"
)

func TestHasPrefix(t *testing.T) {
	muststrings.HasPrefix("hello", "he")

	tests.ExpectPanic(t, func() {
		muststrings.HasPrefix("hello", "hi")
	})
}

func TestHasSuffix(t *testing.T) {
	muststrings.HasSuffix("world", "ld")

	tests.ExpectPanic(t, func() {
		muststrings.HasSuffix("world", "lo")
	})
}
