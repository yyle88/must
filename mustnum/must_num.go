// Package mustnum provides numeric-specific assertion utilities with panic-on-failure semantics
// Implements type-safe validation functions with numeric comparison and state checking
// Supports numeric types spanning integers and floating-points through generic Num constraint
// Integrates with zap structured logging to provide detailed context when assertions are not met
//
// mustnum 提供数值特定的断言工具，带 panic-on-failure 语义
// 实现类型安全的数值比较和状态检查验证函数
// 通过泛型 Num 约束支持所有整数和浮点类型
// 与 zap 结构化日志集成，当断言不满足时提供详细上下文
package mustnum

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Num defines the constraint spanning numeric types including integers and floats
// Num 定义所有数值类型的约束，包括整数和浮点数
type Num interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

// Less validates that a is less than b. Panics if a >= b.
// Less 验证 a 小于 b。如果 a >= b 则触发 panic。
func Less[V Num](a, b V) {
	if a >= b {
		zaplog.ZAPS.Skip1.LOG.Panic("NOT LESS THAN(SHOULD BE LESS)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Lt validates that a is less than b. Alias of Less function. Panics if a >= b.
// Lt 验证 a 小于 b。Less 函数的别名。如果 a >= b 则触发 panic。
func Lt[V Num](a, b V) {
	if a >= b {
		zaplog.ZAPS.Skip1.LOG.Panic("NOT LESS THAN(SHOULD BE LESS)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Lte validates that a is less than / at most b. Panics if a > b.
// Lte 验证 a 小于或等于 b。如果 a > b 则触发 panic。
func Lte[V Num](a, b V) {
	if a > b {
		zaplog.ZAPS.Skip1.LOG.Panic("GREATER THAN(SHOULD BE LESS OR SAME)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Gt validates that a exceeds b. Panics if a <= b.
// Gt 验证 a 大于 b。如果 a <= b 则触发 panic。
func Gt[V Num](a, b V) {
	if a <= b {
		zaplog.ZAPS.Skip1.LOG.Panic("NOT GREATER THAN(SHOULD BE GREATER)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Gte validates that a exceeds / matches b. Panics if a < b.
// Gte 验证 a 大于或等于 b。如果 a < b 则触发 panic。
func Gte[V Num](a, b V) {
	if a < b {
		zaplog.ZAPS.Skip1.LOG.Panic("LESS THAN(SHOULD BE GREATER OR SAME)", zap.Any("a", a), zap.Any("b", b))
	}
}

// Nice validates that numeric value is non-zero. Returns the value if non-zero, panics if zero.
// Nice 验证数值非零。如果非零则返回该值，如果为零则触发 panic。
func Nice[V Num](a V) V {
	if a == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS ZERO(SHOULD BE NON-ZERO)", zap.Any("a", a))
	}
	return a
}

// Zero validates that numeric value is precise zero. Panics if non-zero.
// Zero 验证数值恰好为零。如果非零则触发 panic。
func Zero[V Num](a V) {
	if a != 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS NOT ZERO(SHOULD BE ZERO)", zap.Any("a", a))
	}
}

// Positive validates that value exceeds zero. Panics if value <= 0.
// Positive 验证值严格大于零。如果值 <= 0 则触发 panic。
func Positive[V Num](v V) {
	if v <= 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("NOT POSITIVE(SHOULD BE POSITIVE)", zap.Any("v", v))
	}
}

// Negative validates that value is below zero. Panics if value >= 0.
// Negative 验证值严格小于零。如果值 >= 0 则触发 panic。
func Negative[V Num](v V) {
	if v >= 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("NOT NEGATIVE(SHOULD BE NEGATIVE)", zap.Any("v", v))
	}
}
