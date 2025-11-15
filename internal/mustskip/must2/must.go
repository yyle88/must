// Package must2 provides assertion functions with Skip2 stack frame adjustment
// Implements panic-on-failure validation with 2-depth skip to produce accurate stack traces
// Used when assertions are wrapped by one extra function invocation
// Integrates with zap logging using Skip2 configuration to report correct source location
//
// must2 提供带 Skip2 栈帧调整的断言函数
// 实现带 2 层跳过的 panic-on-failure 验证，以生成准确的堆栈跟踪
// 当断言被一个额外的函数调用包装时使用
// 与 zap 日志集成，使用 Skip2 配置报告正确的源位置
package must2

import (
	"github.com/yyle88/must/internal/utils"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// True validates the value is true with Skip2 stack frame adjustment. Panics if false.
// True 使用 Skip2 栈帧调整验证值为 true。如果为 false 则触发 panic。
func True(v bool) {
	if !v {
		zaplog.ZAPS.Skip2.LOG.Panic("VALUE IS FALSE(SHOULD BE TRUE)", zap.Bool("v", v))
	}
}

// Done validates no error with Skip2 stack frame adjustment. Panics if error is non-nil.
// Done 使用 Skip2 栈帧调整验证没有错误。如果错误非 nil 则触发 panic。
func Done(err error) {
	if err != nil {
		zaplog.ZAPS.Skip2.LOG.Panic("EXPECTED NO ERROR(BUT HAS ERROR)", zap.Error(err))
	}
}

// Nice validates non-zero value with Skip2 stack frame adjustment. Returns value if non-zero, panics if zero.
// Nice 使用 Skip2 栈帧调整验证非零值。如果非零则返回值，如果为零则触发 panic。
func Nice[V comparable](a V) V {
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip2.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)", zap.Any("a", a))
	}
	return a
}

// Same validates values match with Skip2 stack frame adjustment. Panics if not matching.
// Same 使用 Skip2 栈帧调整验证值相等。如果不相等则触发 panic。
func Same[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip2.LOG.Panic("VALUES NOT SAME(SHOULD BE SAME)", zap.Any("a", a), zap.Any("b", b))
	}
}
