package must

import (
	"github.com/pkg/errors"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// True expects the value to be true. Panics if the value is false.
// True 期望值为 true。如果值为 false，则触发 panic。
func True(v bool) {
	if !v {
		zaplog.ZAPS.P1.LOG.Panic("expect TRUE while got FALSE", zap.Bool("v", v))
	}
}

// Done expects no error. Panics if the provided error is non-nil.
// Done 期望没有错误。如果提供的错误不为 nil，则触发 panic。
func Done(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

// Nice expects a non-zero value. Panics if the value is zero, returns the value if non-zero.
// Nice 期望一个非零值。如果值为零，则触发 panic；如果值非零，则返回该值。
func Nice[V comparable](a V) V {
	var b V // zero value
	if a == b {
		zaplog.ZAPS.P1.LOG.Panic("A IS ZERO VALUE", zap.Any("a", a))
	}
	return a
}

// Zero expects a zero value. Panics if the value is non-zero.
// Zero 期望值为零。如果值不为零，则触发 panic。
func Zero[V comparable](a V) {
	var b V // zero value
	if a != b {
		zaplog.ZAPS.P1.LOG.Panic("A IS NOT ZERO VALUE", zap.Any("a", a))
	}
}

// None expects a zero value (empty). Panics if the value is non-zero.
// None 期望值为零（空）。如果值不为零，则触发 panic。
func None[V comparable](a V) {
	var b V // zero value
	if a != b {
		zaplog.ZAPS.P1.LOG.Panic("A IS NOT NONE VALUE", zap.Any("a", a))
	}
}

// Equals expects the values to be equal. Panics if they are not equal.
// Equals 期望值相等。如果值不相等，则触发 panic。
func Equals[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.P1.LOG.Panic("A AND B ARE NOT EQUALS", zap.Any("a", a), zap.Any("b", b))
	}
}

// Is expects equality, not the logic of errors.Is, but the logic of Equals. Panics if the values are not equal.
// Is 期望相等，不是 errors.Is 的逻辑，而是 Equals 的逻辑。如果值不相等，则触发 panic。
func Is[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.P1.LOG.Panic("A AND B ARE NOT EQUALS", zap.Any("a", a), zap.Any("b", b))
	}
}

// Ise expects the errors to be equal, similar to the behavior of errors.Is. Panics if they are not equal.
// Ise 期望错误相等，类似于 errors.Is 的行为。如果错误不相等，则触发 panic。
func Ise(err, target error) {
	if !errors.Is(err, target) {
		zaplog.ZAPS.P1.LOG.Panic("ERROR IS NOT SAME", zap.Error(err), zap.Error(target))
	}
}

// Ok expects a non-zero value. Panics if the value is zero.
// Ok 期望一个非零值。如果值为零，则触发 panic。
func Ok[V comparable](a V) {
	var b V // zero value
	if a == b {
		zaplog.ZAPS.P1.LOG.Panic("A IS ZERO VALUE NOT OK", zap.Any("a", a))
	}
}

// OK expects a non-zero value. Panics if the value is zero. Provides an alternative name for preference.
// OK 期望一个非零值。如果值为零，则触发 panic。提供一个偏好的替代名称。
func OK[V comparable](a V) {
	var b V // zero value
	if a == b {
		zaplog.ZAPS.P1.LOG.Panic("A IS ZERO VALUE NOT OK", zap.Any("a", a))
	}
}

// TRUE expects the value to be true. Panics if the value is false.
// TRUE 期望值为 true。如果值为 false，则触发 panic。
func TRUE(v bool) {
	if !v {
		zaplog.ZAPS.P1.LOG.Panic("expect TRUE while got FALSE", zap.Bool("v", v))
	}
}

// FALSE expects the value to be false. Panics if the value is true.
// FALSE 期望值为 false。如果值为 true，则触发 panic。
func FALSE(v bool) {
	if v {
		zaplog.ZAPS.P1.LOG.Panic("expect FALSE while got TRUE", zap.Bool("v", v))
	}
}

// False expects the value to be false. Panics if the value is true.
// False 期望值为 false。如果值为 true，则触发 panic。
func False(v bool) {
	if v {
		zaplog.ZAPS.P1.LOG.Panic("expect FALSE while got TRUE", zap.Bool("v", v))
	}
}

// Have checks that the slice is not empty. Panics if the slice is empty.
// Have 检查切片是否为空。如果切片为空，则触发 panic。
func Have[T any](a []T) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an empty slice")
	}
}

// Length expects the slice to have length n. Panics if the length is not n.
// Length 期望切片的长度为 n。如果长度不是 n，则触发 panic。
func Length[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len is an abbreviation of Length, serving the same purpose. Panics if the length is not n.
// Len 是 Length 的缩写，功能相同。如果长度不是 n，则触发 panic。
func Len[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// In checks if the value is in the slice. Panics if the value is not found.
// In 检查值是否在切片中。如果未找到该值，则触发 panic。
func In[T comparable](v T, a []T) {
	for i := range a {
		if a[i] == v {
			return
		}
	}
	zaplog.ZAPS.P1.LOG.Panic("expect value in slice while not in", zap.Any("v", v), zap.Int("len", len(a)))
}

// Contains checks if the slice contains the value. Panics if the value is not found.
// Contains 检查切片是否包含该值。如果未找到该值，则触发 panic。
func Contains[T comparable](a []T, v T) {
	for i := range a {
		if a[i] == v {
			return
		}
	}
	zaplog.ZAPS.P1.LOG.Panic("expect slice contains value while not contains", zap.Int("len", len(a)), zap.Any("v", v))
}
