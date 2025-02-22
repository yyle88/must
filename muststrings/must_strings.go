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

// NotHasPrefix checks if the string does not have the specified prefix, panics if it does.
// NotHasPrefix 检查字符串是否没有指定的前缀，有则触发 panic。
func NotHasPrefix(a string, prefix string) {
	if strings.HasPrefix(a, prefix) {
		zaplog.ZAPS.Skip1.LOG.Panic("expect string to not have prefix while it does", zap.String("string", a), zap.String("prefix", prefix))
	}
}

// NotHasSuffix checks if the string does not have the specified suffix, panics if it does.
// NotHasSuffix 检查字符串是否没有指定的后缀，有则触发 panic。
func NotHasSuffix(a string, suffix string) {
	if strings.HasSuffix(a, suffix) {
		zaplog.ZAPS.Skip1.LOG.Panic("expect string to not have suffix while it does", zap.String("string", a), zap.String("suffix", suffix))
	}
}

// Contains checks if the string contains the specified substring, panics if not.
// Contains 检查字符串是否包含指定的子串，没有则触发 panic。
func Contains(a string, sub string) {
	if !strings.Contains(a, sub) {
		zaplog.ZAPS.Skip1.LOG.Panic("expect string to contain substring while it does not", zap.String("string", a), zap.String("substring", sub))
	}
}

// NotContains checks if the string does not contain the specified substring, panics if it does.
// NotContains 检查字符串是否不包含指定的子串，有则触发 panic。
func NotContains(a string, sub string) {
	if strings.Contains(a, sub) {
		zaplog.ZAPS.Skip1.LOG.Panic("expect string to not contain substring while it does", zap.String("string", a), zap.String("substring", sub))
	}
}
