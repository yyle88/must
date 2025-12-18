// Package mustslice provides slice-specific assertion utilities with panic-on-failure semantics
// Implements type-safe slice validation functions using Go generics and standard slices package
// Supports comparison operations, membership testing, and length validation on slice types
// Integrates with zap structured logging to provide detailed context when assertions are not met
//
// mustslice 提供切片特定的断言工具，带 panic-on-failure 语义
// 使用 Go 泛型和标准 slices 包实现类型安全的切片验证函数
// 支持切片类型的比较操作、成员测试和长度验证
// 与 zap 结构化日志集成，当断言不满足时提供详细上下文
package mustslice

import (
	"slices"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Equals checks if two slices match, panics if not.
// Equals 检查两个切片是否相等，不相等则触发 panic。
func Equals[V comparable](a, b []V) {
	if !slices.Equal(a, b) {
		zaplog.ZAPS.Skip1.LOG.Panic("NOT SAME(SHOULD BE SAME)", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// Diff checks if two slices are distinct, panics if matching.
// Diff 检查两个切片是否不同，如果相等则触发 panic。
func Diff[V comparable](a, b []V) {
	if slices.Equal(a, b) {
		zaplog.ZAPS.Skip1.LOG.Panic("ARE SAME(SHOULD BE DIFFERENT)", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// Different checks if two slices are distinct, panics if matching.
// Different 检查两个切片是否不同，如果相等则触发 panic。
func Different[V comparable](a, b []V) {
	if slices.Equal(a, b) {
		zaplog.ZAPS.Skip1.LOG.Panic("ARE SAME(SHOULD BE DIFFERENT)", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// In checks if an element exists in a slice, panics if not.
// In 检查某个元素是否存在于切片中，不存在则触发 panic。
func In[T comparable](v T, a []T) {
	if !slices.Contains(a, v) {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE NOT IN SLICE(SHOULD BE IN)", zap.Any("v", v), zap.Int("len", len(a)))
	}
}

// Contains checks if a slice contains a specific element, panics if not.
// Contains 检查切片是否包含某个特定元素，不包含则触发 panic。
func Contains[T comparable](a []T, v T) {
	if !slices.Contains(a, v) {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE NOT IN SLICE(SHOULD BE IN)", zap.Int("len", len(a)), zap.Any("v", v))
	}
}

// Have function ensures the slice is not vacant, panics if it is.
// Have 确保切片不为空，为空则触发 panic。
func Have[T any](a []T) []T {
	if len(a) == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("SLICE IS EMPTY(SHOULD HAVE ITEMS)")
	}
	return a
}

// Nice ensures the slice is not vacant, and returns it if it contains elements.
// Nice 确保切片不为空，若切片有元素则返回它。
func Nice[T any](a []T) []T {
	if len(a) == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("SLICE IS EMPTY(SHOULD HAVE ITEMS)")
	}
	return a
}

// Zero ensures the slice is vacant, panics if contains elements.
// Zero 确保切片为空，若有元素则触发 panic。
func Zero[T any](a []T) {
	if len(a) != 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("SLICE NOT EMPTY(SHOULD BE EMPTY)")
	}
}

// None ensures the slice is vacant, panics if not.
// None 确保切片内容为空，若有元素则 panic。
func None[T any](a []T) {
	if len(a) != 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("SLICE NOT EMPTY(SHOULD BE EMPTY)")
	}
}

// Length checks if the slice's length matches the expected value, panics if not.
// Length 检查切片的长度是否等于期望值，不等则触发 panic。
func Length[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len checks if the slice's length matches the expected value, panics if not.
// Len 检查切片的长度是否等于期望值，不等则触发 panic。
func Len[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
	}
}
