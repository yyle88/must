package muststrings_test

import (
	"testing"

	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/muststrings"
)

func TestLength(t *testing.T) {
	muststrings.Length("abc", 3)

	tests.ExpectPanic(t, func() {
		muststrings.Length("xyz", 0)
	})
}

func TestLen(t *testing.T) {
	muststrings.Len("123", 3)

	tests.ExpectPanic(t, func() {
		muststrings.Len("000", 0)
	})
}

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
