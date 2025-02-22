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

func TestNotHasPrefix(t *testing.T) {
	muststrings.NotHasPrefix("hello", "hi")

	tests.ExpectPanic(t, func() {
		muststrings.NotHasPrefix("hello", "he")
	})
}

func TestNotHasSuffix(t *testing.T) {
	muststrings.NotHasSuffix("world", "lo")

	tests.ExpectPanic(t, func() {
		muststrings.NotHasSuffix("world", "ld")
	})
}

func TestContains(t *testing.T) {
	muststrings.Contains("hello world", "world")

	tests.ExpectPanic(t, func() {
		muststrings.Contains("hello world", "planet")
	})
}

func TestNotContains(t *testing.T) {
	muststrings.NotContains("hello world", "planet")

	tests.ExpectPanic(t, func() {
		muststrings.NotContains("hello world", "world")
	})
}
