package must

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// True 期望真，当传入假时 panic
func True(v bool) {
	if !v {
		zaplog.ZAPS.P1.LOG.Panic("expect TRUE while got FALSE", zap.Bool("v", v))
	}
}

// Done 期望无错误，假如传入的错误非空时 panic，认为起名为 NoError 也是不错的选择，但实际不要这样用
func Done(err error) {
	if err != nil {
		zaplog.ZAPS.P1.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

// Nice 期望非零值，当传入零值时 panic
func Nice[V comparable](a V) {
	var b V //zero
	if a == b {
		zaplog.ZAPS.P1.LOG.Panic("A IS ZERO VALUE", zap.Any("a", a))
	}
}

// Zero 期望零值，当传入非零值时 panic
func Zero[V comparable](a V) {
	var b V //zero
	if a != b {
		zaplog.ZAPS.P1.LOG.Panic("A IS NOT ZERO VALUE", zap.Any("a", a))
	}
}

// Equals 期望相等，假如不相等就 panic
func Equals[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.P1.LOG.Panic("A AND B ARE NOT EQUALS", zap.Any("a", a), zap.Any("b", b))
	}
}

// Is 期望相等，注意这里不是 errors.Is 的逻辑，而是期望 Equals 的逻辑
func Is[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.P1.LOG.Panic("A AND B ARE NOT EQUALS", zap.Any("a", a), zap.Any("b", b))
	}
}

// Ok 跟 Nice 相同，期望非零值，假如传入零值就 panic
func Ok[V comparable](a V) {
	var b V //zero
	if a == b {
		zaplog.ZAPS.P1.LOG.Panic("A IS ZERO VALUE NOT OK", zap.Any("a", a))
	}
}

// OK 跟 Ok 相同，只是很多时候我们喜欢大写的 OK 而不是这种 Ok，因此两者都提供让调用者自己选择喜欢的
func OK[V comparable](a V) {
	var b V //zero
	if a == b {
		zaplog.ZAPS.P1.LOG.Panic("A IS ZERO VALUE NOT OK", zap.Any("a", a))
	}
}

// TRUE 和 True 逻辑相同，毕竟是断言，直接用大写表达更强硬的意图
func TRUE(v bool) {
	if !v {
		zaplog.ZAPS.P1.LOG.Panic("expect TRUE while got FALSE", zap.Bool("v", v))
	}
}

// FALSE 和 False 逻辑相同，毕竟是断言，直接用大写表达更强硬的意图
func FALSE(v bool) {
	if v {
		zaplog.ZAPS.P1.LOG.Panic("expect FALSE while got TRUE", zap.Bool("v", v))
	}
}

// False 期望假，当传入真的时候 panic，因为很少有期望假的时候，因此这个函数不和 True 放在一起，而是放在末席
func False(v bool) {
	if v {
		zaplog.ZAPS.P1.LOG.Panic("expect FALSE while got TRUE", zap.Bool("v", v))
	}
}

// Have 意思是 NotEmpty 非空 slice
func Have[T any](a []T) {
	if len(a) == 0 {
		zaplog.ZAPS.P1.LOG.Panic("expect LENGTH > 0 while got an empty slice")
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

// In 检查元素是否在数组里
func In[T comparable](v T, a []T) {
	for i := range a {
		if a[i] == v {
			return
		}
	}
	zaplog.ZAPS.P1.LOG.Panic("expect value in slice while not in", zap.Any("v", v), zap.Int("len", len(a)))
}

// Contains 检查数组是否包含元素
func Contains[T comparable](a []T, v T) {
	for i := range a {
		if a[i] == v {
			return
		}
	}
	zaplog.ZAPS.P1.LOG.Panic("expect slice contains value while not contains", zap.Int("len", len(a)), zap.Any("v", v))
}
