// Package mustmap provides map-specific assertion utilities with panic-on-failure semantics
// Implements type-safe map validation functions using Go generics and standard maps package
// Supports comparison operations, length validation, and element existence checking on map types
// Integrates with zap structured logging to provide detailed context when assertions are not met
//
// mustmap 提供 map 特定的断言工具，带 panic-on-failure 语义
// 使用 Go 泛型和标准 maps 包实现类型安全的 map 验证函数
// 支持 map 类型的比较操作、长度验证和元素存在性检查
// 与 zap 结构化日志集成，当断言不满足时提供详细上下文
package mustmap

import (
	"maps"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Equals compares two maps matching. If not matching, it panics.
// Equals 比较两个 map 是否相等，如果不相等，则触发 panic。
func Equals[K, V comparable](a, b map[K]V) {
	if !maps.Equal(a, b) {
		zaplog.ZAPS.Skip1.LOG.Panic("NOT SAME(SHOULD BE SAME)", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// Diff compares two maps to ensure distinction. If matching, it panics.
// Diff 比较两个 map 是否不相等，如果相等，则触发 panic。
func Diff[K, V comparable](a, b map[K]V) {
	if maps.Equal(a, b) {
		zaplog.ZAPS.Skip1.LOG.Panic("ARE SAME(SHOULD BE DIFFERENT)", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// Different compares two maps to ensure distinction. If matching, it panics.
// Different 比较两个 map 是否不相等，如果相等，则触发 panic。
func Different[K, V comparable](a, b map[K]V) {
	if maps.Equal(a, b) {
		zaplog.ZAPS.Skip1.LOG.Panic("ARE SAME(SHOULD BE DIFFERENT)", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// Have checks if a map is non-vacant. If vacant, it panics.
// Have 检查一个 map 是否非空，如果为空，则触发 panic。
func Have[K comparable, V any](a map[K]V) map[K]V {
	if len(a) == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("MAP IS EMPTY(SHOULD HAVE ITEMS)")
	}
	return a
}

// Nice checks if a map is non-vacant and returns it. If vacant, it panics.
// Nice 检查一个 map 是否非空并返回它，如果为空，则触发 panic。
func Nice[K comparable, V any](a map[K]V) map[K]V {
	if len(a) == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("MAP IS EMPTY(SHOULD HAVE ITEMS)")
	}
	return a
}

// Zero ensures the map is vacant, panics if contains entries.
// Zero 确保 map 为空，若有条目则触发 panic。
func Zero[K comparable, V any](a map[K]V) {
	if len(a) != 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("MAP NOT EMPTY(SHOULD BE EMPTY)")
	}
}

// None ensures the map is vacant, panics if not.
// None 确保 map 内容为空，若有元素则 panic。
func None[K comparable, V any](a map[K]V) {
	if len(a) != 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("MAP NOT EMPTY(SHOULD BE EMPTY)")
	}
}

// Length checks if the length of a map matches n. If not, it panics.
// Length 检查一个 map 的长度是否等于 n，如果不等，则触发 panic。
func Length[K comparable, V any](a map[K]V, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len checks if the length of a map matches n. If not, it panics.
// Len 是 Length 的简写版本，检查一个 map 的长度是否等于 n，如果不等，则触发 panic。
func Len[K comparable, V any](a map[K]V, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Get func get value of element from the map. If the element does not exist, it panics.
// Get 根据给定的键从 map 中检索值，如果键不存在，则触发 panic。
func Get[K, V comparable](a map[K]V, key K) V {
	value, exists := a[key]
	if !exists {
		zaplog.ZAPS.Skip1.LOG.Panic("KEY NOT IN MAP(SHOULD BE IN)", zap.Any("key", key))
	}
	return value
}
