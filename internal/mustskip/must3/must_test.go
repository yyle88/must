package must3_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/mustskip/must3"
)

func TestTrue(t *testing.T) {
	must3.True(true)

	require.Panics(t, func() {
		must3.True(false)
	})
}

func TestDone(t *testing.T) {
	must3.Done(error(nil))

	require.Panics(t, func() {
		must3.Done(errors.New("wa"))
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, 88, must3.Nice(88))
	require.Equal(t, "xyz", must3.Nice("xyz"))
	require.Equal(t, 3.1415926, must3.Nice(3.1415926))

	require.Panics(t, func() {
		must3.Nice(0.0)
	})
}

func TestSame(t *testing.T) {
	must3.Same("abc", "abc")
	must3.Same(123, 123)
	must3.Same(0.8, 0.8)

	require.Panics(t, func() {
		must3.Same("abc", "xyz")
	})
}
