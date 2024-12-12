package mustmap

import (
	"maps"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Equals compares two maps for equality. If they are not equal, it panics.
// Equals 比较两个 map 是否相等，如果不相等，则触发 panic。
func Equals[K, V comparable](a, b map[K]V) {
	if !maps.Equal(a, b) {
		zaplog.ZAPS.P1.LOG.Panic("expect map equals while not equals", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// Have checks if a map is non-empty. If it is empty, it panics.
// Have 检查一个 map 是否非空，如果为空，则触发 panic。
func Have[K comparable, V any](a map[K]V) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none map")
	}
}

// Nice checks if a map is non-empty and returns it. If it is empty, it panics.
// Nice 检查一个 map 是否非空并返回它，如果为空，则触发 panic。
func Nice[K comparable, V any](a map[K]V) map[K]V {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none map")
	}
	return a
}

// None ensures the map is empty, panics if not.
// None 确保 map 内容为空，若有元素则 panic。
func None[K comparable, V any](a map[K]V) {
	if len(a) != 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = 0 while contains elements")
	}
}

// Length checks if the length of a map is equal to n. If not, it panics.
// Length 检查一个 map 的长度是否等于 n，如果不等，则触发 panic。
func Length[K comparable, V any](a map[K]V, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len checks if the length of a map is equal to n. If not, it panics.
// Len 是 Length 的简写版本，检查一个 map 的长度是否等于 n，如果不等，则触发 panic。
func Len[K comparable, V any](a map[K]V, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Get func get value of key from the map. If the key does not exist, it panics.
// Get 根据给定的键从 map 中检索值，如果键不存在，则触发 panic。
func Get[K, V comparable](a map[K]V, key K) V {
	value, exists := a[key]
	if !exists {
		zaplog.ZAPS.P1.LOG.Panic("expect KEY to EXIST in map", zap.Any("key", key))
	}
	return value
}
