package mustboolean

import (
	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// True expects the value to be true. Panics if the value is false.
// True 期望值为 true。如果值为 false，则触发 panic。
func True(v bool) {
	if !v {
		zaplog.ZAPS.Skip1.LOG.Panic("VALUE IS FALSE(SHOULD BE TRUE)", zap.Bool("v", v))
	}
}

// Conflict ensures at most one boolean is true. Panics if multiple are true.
// Conflict 确保最多只有一个布尔值为 true。如果有多个为 true，则触发 panic。
func Conflict(bs ...bool) {
	firstIndex := -1
	for idx, b := range bs {
		if b {
			if firstIndex >= 0 {
				zaplog.ZAPS.Skip1.LOG.Panic("conflict: multiple true values", zap.Int("first", firstIndex), zap.Int("second", idx))
			}
			firstIndex = idx
		}
	}
}
