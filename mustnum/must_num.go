package mustnum

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

type Num interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

func Less[V Num](a, b V) {
	if a >= b {
		zaplog.ZAPS.Skip1.LOG.Panic("expect Less while not", zap.Any("a", a), zap.Any("b", b))
	}
}

func Lt[V Num](a, b V) {
	if a >= b {
		zaplog.ZAPS.Skip1.LOG.Panic("expect Lt while not", zap.Any("a", a), zap.Any("b", b))
	}
}

func Lte[V Num](a, b V) {
	if a > b {
		zaplog.ZAPS.Skip1.LOG.Panic("expect Lte while not", zap.Any("a", a), zap.Any("b", b))
	}
}

func Gt[V Num](a, b V) {
	if a <= b {
		zaplog.ZAPS.Skip1.LOG.Panic("expect Gt while not", zap.Any("a", a), zap.Any("b", b))
	}
}

func Gte[V Num](a, b V) {
	if a < b {
		zaplog.ZAPS.Skip1.LOG.Panic("expect Gte while not", zap.Any("a", a), zap.Any("b", b))
	}
}

func Nice[V Num](a V) V {
	if a == 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("A IS ZERO VALUE", zap.Any("a", a))
	}
	return a
}

func Zero[V Num](a V) {
	if a != 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("A IS NOT ZERO VALUE", zap.Any("a", a))
	}
}

func Positive[V Num](v V) {
	if v <= 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("expect POSITIVE while not", zap.Any("v", v))
	}
}

func Negative[V Num](v V) {
	if v >= 0 {
		zaplog.ZAPS.Skip1.LOG.Panic("expect NEGATIVE while not", zap.Any("v", v))
	}
}
