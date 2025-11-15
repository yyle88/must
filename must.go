// Package must provides panic-on-failure assertion utilities with structured logging
// Implements type-safe validation functions that panic with detailed context when expectations are not met
// Supports generic types enabling flexible value checking across different data types
// Integrates with zap logging to provide precise stack trace information
//
// must 提供带结构化日志的 panic-on-failure 断言工具
// 实现类型安全的验证函数，当期望不满足时 panic 并提供详细上下文
// 支持泛型类型，在不同数据类型间灵活检查值
// 与 zap 日志集成，提供精确的堆栈跟踪信息
package must

import (
	"github.com/pkg/errors"
	"github.com/yyle88/must/internal/mustskip/must2"
	"github.com/yyle88/must/internal/utils"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// True expects the value to be true. Panics if the value is false.
// True 期望值为 true。如果值为 false，则触发 panic。
func True(v bool) {
	if !v {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS FALSE(SHOULD BE TRUE)", zap.Bool("v", v))
	}
}

// Done expects no error. Panics if the provided error is non-nil.
// Done 期望没有错误。如果提供的错误不为 nil，则触发 panic。
func Done(err error) {
	if err != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("EXPECTED NO ERROR(BUT HAS ERROR)", zap.Error(err))
	}
}

// Must expects no error. Panics if the provided error is non-nil.
// Must 期望没有错误。如果提供的错误不为 nil，则触发 panic。
func Must(err error) {
	if err != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("HAS ERROR(SHOULD BE NO ERROR)", zap.Error(err))
	}
}

// Nice expects a non-zero value. Panics if the value is zero, returns the value if non-zero.
// Nice 期望一个非零值。如果值为零，则触发 panic；如果值非零，则返回该值。
func Nice[V comparable](a V) V {
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)", zap.Any("a", a))
	}
	return a
}

// Zero expects a zero value. Panics if the value is non-zero.
// Zero 期望值为零。如果值不为零，则触发 panic。
func Zero[V comparable](a V) {
	if a != utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS NOT ZERO(SHOULD BE ZERO)", zap.Any("a", a))
	}
}

// None expects a zero value (vacant/absent). Panics if the value is non-zero.
// None 期望值为零（空）。如果值不为零，则触发 panic。
func None[V comparable](a V) {
	if a != utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS NOT ZERO(SHOULD BE ZERO)", zap.Any("a", a))
	}
}

// Null expects the value to be nil. Panics if the value is non-nil.
// Null 期望值为 nil。如果值不为 nil，则触发 panic。
func Null[T any](v *T) {
	if v != nil {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE PRESENT(SHOULD BE ABSENT)")
	}
}

// Full expects the value to be non-nil. Panics if the value is nil.
// Full 期望值为非 nil。如果值为 nil，则触发 panic。
func Full[T any](v *T) *T {
	if v == nil {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE ABSENT(SHOULD BE PRESENT)")
	}
	return v
}

// Equals expects the values to be the same. Panics if not the same.
// Equals 期望值相等。如果值不相等，则触发 panic。
func Equals[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUES NOT SAME(SHOULD BE SAME)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Same expects the values to be the same. Panics if not the same.
// Same 期望值相等。如果值不相等，则触发 panic。
func Same[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUES NOT SAME(SHOULD BE SAME)", zap.Any("a", a), zap.Any("b", b))
	}
}

// SameNice expects the values to match and be non-zero. Panics if not matching / when zero. Returns the value when conditions are met.
// SameNice 期望值相等且非零。如果值不相等/为零，则触发 panic。如果条件满足，则返回该值。
func SameNice[V comparable](a, b V) V {
	must2.Nice(a)
	must2.Nice(b)
	must2.Same(a, b)
	return a
}

// Sane means same && nice
// Sane 期望值相等且非零。如果值不相等/为零，则触发 panic。如果条件满足，则返回该值。
func Sane[V comparable](a, b V) V {
	must2.Nice(a)
	must2.Nice(b)
	must2.Same(a, b)
	return a
}

// Diff expects the values to be distinct. Panics if the values match.
// Diff 期望值不同。如果值相同，则触发 panic。
func Diff[V comparable](a, b V) {
	if a == b {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUES ARE SAME(SHOULD BE DIFFERENT)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Different expects the values to be distinct. Panics if the values match.
// Different 期望值不同。如果值相同，则触发 panic。
func Different[V comparable](a, b V) {
	if a == b {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUES ARE SAME(SHOULD BE DIFFERENT)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Is expects matching values, not the logic of errors.Is, but the logic of Equals. Panics if the values do not match.
// Is 期望相等，不是 errors.Is 的逻辑，而是 Equals 的逻辑。如果值不相等，则触发 panic。
func Is[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUES NOT SAME(SHOULD BE SAME)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Ise expects the errors to match, using the logic of errors.Is. Panics if not matching.
// Ise 期望错误相等，类似于 errors.Is 的行为。如果错误不相等，则触发 panic。
func Ise(err, target error) {
	if !errors.Is(err, target) {
		zaplog.ZAPS.Skip1.LOG.Panic("ERROR MISMATCH(NOT SAME ERROR)", zap.Error(err), zap.Error(target))
	}
}

// Ok expects a non-zero value. Panics if the value is zero.
// Ok 期望一个非零值。如果值为零，则触发 panic。
func Ok[V comparable](a V) {
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)", zap.Any("a", a))
	}
}

// OK expects a non-zero value. Panics if the value is zero. Provides an alternative name based on preference.
// OK 期望一个非零值。如果值为零，则触发 panic。提供一个偏好的替代名称。
func OK[V comparable](a V) {
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)", zap.Any("a", a))
	}
}

// TRUE expects the value to be true. Panics if the value is false.
// TRUE 期望值为 true。如果值为 false，则触发 panic。
func TRUE(v bool) {
	if !v {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS FALSE(SHOULD BE TRUE)", zap.Bool("v", v))
	}
}

// FALSE expects the value to be false. Panics if the value is true.
// FALSE 期望值为 false。如果值为 true，则触发 panic。
func FALSE(v bool) {
	if v {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS TRUE(SHOULD BE FALSE)", zap.Bool("v", v))
	}
}

// False expects the value to be false. Panics if the value is true.
// False 期望值为 false。如果值为 true，则触发 panic。
func False(v bool) {
	if v {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS TRUE(SHOULD BE FALSE)", zap.Bool("v", v))
	}
}

// Have checks that the slice is not vacant. Panics if the slice is vacant.
// Have 检查切片是否为空。如果切片为空，则触发 panic。
func Have[T any](a []T) []T {
	if len(a) == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("SLICE IS EMPTY(SHOULD HAVE ITEMS)")
	}
	return a
}

// Length expects the slice to have length n. Panics if the length is not n.
// Length 期望切片的长度为 n。如果长度不是 n，则触发 panic。
func Length[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len is an abbreviation of Length, serving the same purpose. Panics if the length is not n.
// Len 是 Length 的缩写，功能相同。如果长度不是 n，则触发 panic。
func Len[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
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
	zaplog.ZAPS.Skip1.LOG.Panic("VALUE NOT IN SLICE(SHOULD BE IN)", zap.Any("v", v), zap.Int("len", len(a)))
}

// Contains checks if the slice contains the value. Panics if the value is not found.
// Contains 检查切片是否包含该值。如果未找到该值，则触发 panic。
func Contains[T comparable](a []T, v T) {
	for i := range a {
		if a[i] == v {
			return
		}
	}
	zaplog.ZAPS.Skip1.LOG.Panic("VALUE NOT IN SLICE(SHOULD BE IN)", zap.Int("len", len(a)), zap.Any("v", v))
}
