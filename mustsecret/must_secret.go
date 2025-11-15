// Package mustsecret provides assertion utilities with panic-on-failure semantics designed to protect sensitive data
// Implements validation functions that avoid logging data values to prevent information leakage
// Panic messages exclude data values to prevent leaking secrets in logs
// Integrates with zap structured logging but omits sensitive value details
//
// mustsecret 提供断言工具，带 panic-on-failure 语义，专门用于保护敏感数据
// 实现避免记录实际值的验证函数，以防止信息泄露
// 所有 panic 消息都排除数据值，防止在日志中泄露机密
// 与 zap 结构化日志集成，但省略敏感值详情
package mustsecret

import (
	"github.com/yyle88/must/internal/utils"
	"github.com/yyle88/zaplog"
)

// Nice expects a non-zero value. Panics if the value is zero, returns the value if non-zero.
// Nice 期望一个非零值。如果值为零，则触发 panic；如果值非零，则返回该值。
func Nice[V comparable](a V) V {
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)") // not show data in the log message
	}
	return a
}

// Zero expects a zero value. Panics if the value is non-zero.
// Zero 期望值为零。如果值不为零，则触发 panic。
func Zero[V comparable](a V) {
	if a != utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS NOT ZERO(SHOULD BE ZERO)") // not show data in the log message
	}
}

// Same expects the values to match. Panics if not matching.
// Same 期望值相等。如果值不相等，则触发 panic。
func Same[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUES NOT SAME(SHOULD BE SAME)") // not show data in the log message
	}
}

// Sane means same && nice
// Sane 期望值相等且非零。如果值不相等/为零，则触发 panic。如果条件满足，则返回该值。
func Sane[V comparable](a, b V) V {
	if a != b {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUES NOT SAME(SHOULD BE SAME)") // not show data in the log message
	}
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)") // not show data in the log message
	}
	return a
}
