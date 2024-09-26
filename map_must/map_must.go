package map_must

import (
	"maps"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Equals[K, V comparable](a, b map[K]V) {
	if !maps.Equal(a, b) {
		zaplog.ZAPS.P1.LOG.Panic("expect map equals while not equals", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// Have 意思是 NotEmpty 非空 map
func Have[K comparable, V any](a map[K]V) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none map")
	}
}

// Nice 意思是 NotEmpty 非空 map
// must.Nice(a) 的作用仅仅是判定是否为空，而这里的作用是判断是否有内容，但这样易混淆，因此建议使用 Have 函数
func Nice[K comparable, V any](a map[K]V) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none map")
	}
}

// Length 期望长度是 n，否则 panic
func Length[K comparable, V any](a map[K]V, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len 和 Length 作用相同，只是缩写，我更倾向于使用缩写
func Len[K comparable, V any](a map[K]V, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}
