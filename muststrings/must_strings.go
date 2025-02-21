package muststrings

import (
	"strings"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// HasPrefix checks if the string has the specified prefix, panics if not.
// HasPrefix 检查字符串是否有指定的前缀，没有则触发 panic。
func HasPrefix(a string, prefix string) {
	if !strings.HasPrefix(a, prefix) {
		zaplog.ZAPS.Skip1.LOG.Panic("expect string to have prefix while it does not", zap.String("string", a), zap.String("prefix", prefix))
	}
}

// HasSuffix checks if the string has the specified suffix, panics if not.
// HasSuffix 检查字符串是否有指定的后缀，没有则触发 panic。
func HasSuffix(a string, suffix string) {
	if !strings.HasSuffix(a, suffix) {
		zaplog.ZAPS.Skip1.LOG.Panic("expect string to have suffix while it does not", zap.String("string", a), zap.String("suffix", suffix))
	}
}
