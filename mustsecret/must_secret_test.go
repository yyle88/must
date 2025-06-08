package mustsecret_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/mustsecret"
)

func TestNice(t *testing.T) {
	require.Equal(t, int8(5), mustsecret.Nice(int8(5)))
	require.Equal(t, "hello", mustsecret.Nice("hello"))
	require.Equal(t, 1.23, mustsecret.Nice(1.23))

	tests.ExpectPanic(t, func() {
		mustsecret.Nice("")
	})
}

func TestZero(t *testing.T) {
	mustsecret.Zero(uint32(0))
	mustsecret.Zero(float32(0))
	mustsecret.Zero("")

	tests.ExpectPanic(t, func() {
		mustsecret.Zero("some-value")
	})
}

func TestSame_NewData(t *testing.T) {
	mustsecret.Same(true, true)
	mustsecret.Same(uint8(255), uint8(255))
	mustsecret.Same("hello", "hello")
	mustsecret.Same(float32(3.14), float32(3.14))

	tests.ExpectPanic(t, func() {
		mustsecret.Same(100, 200)
	})

	tests.ExpectPanic(t, func() {
		mustsecret.Same(false, true)
	})
}

func TestSane(t *testing.T) {
	require.Equal(t, "abc", mustsecret.Sane("abc", "abc"))
	require.Equal(t, 123, mustsecret.Sane(123, 123))
	require.Equal(t, 0.8, mustsecret.Sane(0.8, 0.8))

	tests.ExpectPanic(t, func() {
		mustsecret.Sane("abc", "xyz")
	})

	tests.ExpectPanic(t, func() {
		mustsecret.Sane(0, 0)
	})
}
