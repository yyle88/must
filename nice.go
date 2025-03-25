package must

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func nice[V comparable](a V) V {
	var z V // zero value
	if a == z {
		zaplog.ZAPS.Skip2.LOG.Panic("A IS ZERO VALUE", zap.Any("a", a))
	}
	return a
}

func same[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip2.LOG.Panic("A AND B ARE NOT SAME", zap.Any("a", a), zap.Any("b", b))
	}
}
