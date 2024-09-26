package slice_must

import (
	"slices"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func Equals[V comparable](a, b []V) {
	if !slices.Equal(a, b) {
		zaplog.ZAPS.P1.LOG.Panic("expect slice equals while not equals", zap.Int("len_a", len(a)), zap.Int("len_b", len(b)))
	}
}

// In 检查元素是否在数组里
func In[T comparable](v T, a []T) {
	if !slices.Contains(a, v) {
		zaplog.ZAPS.P1.LOG.Panic("expect value in slice while not in", zap.Any("v", v), zap.Int("len", len(a)))
	}
}

// Contains 检查数组是否包含元素，假如不包含就 panic
func Contains[T comparable](a []T, v T) {
	if !slices.Contains(a, v) {
		zaplog.ZAPS.P1.LOG.Panic("expect slice contains value while not contains", zap.Int("len", len(a)), zap.Any("v", v))
	}
}

// Have 意思是 NotEmpty 非空 slice
func Have[T any](a []T) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none slice")
	}
}

// Nice 意思是 NotEmpty 非空 slice
// must.Nice(a) 的作用仅仅是判定是否为空，而这里的作用是判断是否有内容，但这样易混淆，因此建议使用 Have 函数
func Nice[T any](a []T) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an none slice")
	}
}

// Length 期望长度是 n，否则 panic
func Length[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len 和 Length 作用相同，只是缩写，我更倾向于使用缩写
func Len[T any](a []T, n int) {
	if len(a) != n {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH = n while not equals", zap.Int("len", len(a)), zap.Int("n", n))
	}
}
