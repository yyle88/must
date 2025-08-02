package mustskip2c_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/mustskip/mustskip2c"
)

func TestTrue(t *testing.T) {
	mustskip2c.True(true)

	tests.ExpectPanic(t, func() {
		mustskip2c.True(false)
	})
}

func TestDone(t *testing.T) {
	mustskip2c.Done(error(nil))

	tests.ExpectPanic(t, func() {
		mustskip2c.Done(errors.New("wa"))
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, 88, mustskip2c.Nice(88))
	require.Equal(t, "xyz", mustskip2c.Nice("xyz"))
	require.Equal(t, 3.1415926, mustskip2c.Nice(3.1415926))

	tests.ExpectPanic(t, func() {
		mustskip2c.Nice(0.0)
	})
}

func TestSame(t *testing.T) {
	mustskip2c.Same("abc", "abc")
	mustskip2c.Same(123, 123)
	mustskip2c.Same(0.8, 0.8)

	tests.ExpectPanic(t, func() {
		mustskip2c.Same("abc", "xyz")
	})
}
