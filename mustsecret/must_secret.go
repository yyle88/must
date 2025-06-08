package mustsecret

import (
	"github.com/yyle88/must/internal/utils"
	"github.com/yyle88/zaplog"
)

// Nice expects a non-zero value. Panics if the value is zero, returns the value if non-zero.
// Nice 期望一个非零值。如果值为零，则触发 panic；如果值非零，则返回该值。
func Nice[V comparable](a V) V {
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("A IS ZERO VALUE") // not show data in the log message
	}
	return a
}

// Zero expects a zero value. Panics if the value is non-zero.
// Zero 期望值为零。如果值不为零，则触发 panic。
func Zero[V comparable](a V) {
	if a != utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("A IS NOT ZERO VALUE") // not show data in the log message
	}
}

// Same expects the values to be same. Panics if they are not same.
// Same 期望值相等。如果值不相等，则触发 panic。
func Same[V comparable](a, b V) {
	if a != b {
		zaplog.ZAPS.Skip1.LOG.Panic("A AND B ARE NOT SAME") // not show data in the log message
	}
}

// Sane means same && nice
// Sane 期望值相等且非零。如果值不相等或为零，则触发 panic。如果条件满足，则返回该值。
func Sane[V comparable](a, b V) V {
	if a != b {
		zaplog.ZAPS.Skip1.LOG.Panic("A AND B ARE NOT SAME") // not show data in the log message
	}
	if a == utils.Zero[V]() {
		zaplog.ZAPS.Skip1.LOG.Panic("A IS ZERO VALUE") // not show data in the log message
	}
	return a
}
