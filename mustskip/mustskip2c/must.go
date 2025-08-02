package mustskip2c

import (
	"github.com/yyle88/must/internal/utils"
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

func True(v bool) {
	if !v {
		zaplog.ZAPS.Skip2.LOG.Panic("expect TRUE while got FALSE", zap.Bool("v", v))
	}
}

func Done(err error) {
	if err != nil {
		zaplog.ZAPS.Skip2.LOG.Panic("NO ERROR BUG", zap.Error(err))
	}
}

func Nice[V comparable](a V) V {
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip2.LOG.Panic("A IS ZERO VALUE", zap.Any("a", a))
	}
	return a
}

func Same[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip2.LOG.Panic("A AND B ARE NOT SAME", zap.Any("a", a), zap.Any("b", b))
	}
}
