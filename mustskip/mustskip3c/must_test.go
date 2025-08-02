package mustskip3c_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/mustskip/mustskip3c"
)

func TestTrue(t *testing.T) {
	mustskip3c.True(true)

	tests.ExpectPanic(t, func() {
		mustskip3c.True(false)
	})
}

func TestDone(t *testing.T) {
	mustskip3c.Done(error(nil))

	tests.ExpectPanic(t, func() {
		mustskip3c.Done(errors.New("wa"))
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, 88, mustskip3c.Nice(88))
	require.Equal(t, "xyz", mustskip3c.Nice("xyz"))
	require.Equal(t, 3.1415926, mustskip3c.Nice(3.1415926))

	tests.ExpectPanic(t, func() {
		mustskip3c.Nice(0.0)
	})
}

func TestSame(t *testing.T) {
	mustskip3c.Same("abc", "abc")
	mustskip3c.Same(123, 123)
	mustskip3c.Same(0.8, 0.8)

	tests.ExpectPanic(t, func() {
		mustskip3c.Same("abc", "xyz")
	})
}
