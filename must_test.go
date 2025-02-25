package must_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
	"github.com/yyle88/must/internal/tests"
)

func TestTrue(t *testing.T) {
	must.True(true)

	tests.ExpectPanic(t, func() {
		must.True(false)
	})
}

func TestDone(t *testing.T) {
	must.Done(error(nil))

	tests.ExpectPanic(t, func() {
		must.Done(errors.New("wa"))
	})
}

func TestMust(t *testing.T) {
	must.Must(error(nil))

	tests.ExpectPanic(t, func() {
		must.Must(errors.New("wa"))
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, 88, must.Nice(88))
	require.Equal(t, "xyz", must.Nice("xyz"))
	require.Equal(t, 3.1415926, must.Nice(3.1415926))

	tests.ExpectPanic(t, func() {
		must.Nice(0.0)
	})
}

func TestZero(t *testing.T) {
	must.Zero(0)
	must.Zero("")
	must.Zero(0.0)

	tests.ExpectPanic(t, func() {
		must.Zero(uint64(200))
	})
}

func TestNone(t *testing.T) {
	must.None(0)
	must.None(0.0)
	must.None("")
	must.None(error(nil))

	tests.ExpectPanic(t, func() {
		must.None(errors.New("wa"))
	})
}

func TestEquals(t *testing.T) {
	must.Equals("abc", "abc")
	must.Equals(123, 123)
	must.Equals(0.8, 0.8)

	tests.ExpectPanic(t, func() {
		must.Equals("abc", "xyz")
	})
}

func TestSame(t *testing.T) {
	must.Same("abc", "abc")
	must.Same(123, 123)
	must.Same(0.8, 0.8)

	tests.ExpectPanic(t, func() {
		must.Same("abc", "xyz")
	})
}

func TestDifferent(t *testing.T) {
	must.Different("abc", "xyz")
	must.Different(123, 321)
	must.Different(0.8, 0.88)

	tests.ExpectPanic(t, func() {
		must.Different("abc", "abc")
	})
}

func TestIs(t *testing.T) {
	must.Is(123, 123)
	must.Is("!", "!")
	must.Is('!', '!')
	must.Is(0.5, 0.5)

	tests.ExpectPanic(t, func() {
		must.Is(1, 2)
	})
}

func TestIse(t *testing.T) {
	errA := errors.New("a")
	must.Ise(errA, errA)

	tests.ExpectPanic(t, func() {
		must.Ise(errors.New("a"), errors.New("b"))
	})
}

func TestOk(t *testing.T) {
	must.Ok(1990)

	tests.ExpectPanic(t, func() {
		must.Ok("")
	})
}

func TestOK(t *testing.T) {
	must.OK("yyle")
	must.OK(88)
	must.OK(0.88)
	must.OK(uint64(88))
	must.OK('8')

	tests.ExpectPanic(t, func() {
		must.OK(0)
	})
}

func TestTRUE(t *testing.T) {
	must.TRUE(true)

	tests.ExpectPanic(t, func() {
		must.TRUE(false)
	})
}

func TestFALSE(t *testing.T) {
	must.FALSE(false)

	tests.ExpectPanic(t, func() {
		must.FALSE(true)
	})
}

func TestFalse(t *testing.T) {
	must.False(false)

	tests.ExpectPanic(t, func() {
		must.False(true)
	})
}

func TestHave(t *testing.T) {
	must.Have([]int{0, 1, 2, 3})
	must.Have([]string{"abc", "xyz"})
	must.Have([]float64{0.0})

	tests.ExpectPanic(t, func() {
		must.Have([]uint64{})
	})
}

func TestLength(t *testing.T) {
	must.Length([]int{0, 1, 2, 3}, 4)
	must.Length([]string{"abc", "xyz"}, 2)
	must.Length([]float64{0.0}, 1)
	must.Length([]uint64{}, 0)

	tests.ExpectPanic(t, func() {
		must.Length([]any{"a", 1, 'x'}, 2)
	})
}

func TestLen(t *testing.T) {
	must.Len([]int{}, 0)
	must.Len([]float64{0.1, 0.2}, 2)

	tests.ExpectPanic(t, func() {
		must.Len([]uint64{}, 88)
	})
}

func TestIn(t *testing.T) {
	must.In(1, []int{1, 2, 3})
	must.In(2.0, []int{1, 2.0, 3})

	tests.ExpectPanic(t, func() {
		must.In("a", []string{"b", "c"})
	})
}

func TestContains(t *testing.T) {
	must.Contains([]string{"a", "b", "c"}, "a")

	tests.ExpectPanic(t, func() {
		must.Contains([]int{1, 2, 3}, 4)
	})
}

type Example struct {
	S string
}

func TestNull(t *testing.T) {
	var example *Example
	must.Null(example)

	tests.ExpectPanic(t, func() {
		must.Null(&Example{S: "abc"})
	})
}

func TestFull(t *testing.T) {
	must.Full(&Example{S: "abc"})

	tests.ExpectPanic(t, func() {
		var example *Example
		must.Full(example)
	})
}
