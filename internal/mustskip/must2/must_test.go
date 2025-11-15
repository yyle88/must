package must2_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/mustskip/must2"
)

func TestTrue(t *testing.T) {
	must2.True(true)

	require.Panics(t, func() {
		must2.True(false)
	})
}

func TestDone(t *testing.T) {
	must2.Done(error(nil))

	require.Panics(t, func() {
		must2.Done(errors.New("wa"))
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, 88, must2.Nice(88))
	require.Equal(t, "xyz", must2.Nice("xyz"))
	require.Equal(t, 3.1415926, must2.Nice(3.1415926))

	require.Panics(t, func() {
		must2.Nice(0.0)
	})
}

func TestSame(t *testing.T) {
	must2.Same("abc", "abc")
	must2.Same(123, 123)
	must2.Same(0.8, 0.8)

	require.Panics(t, func() {
		must2.Same("abc", "xyz")
	})
}
