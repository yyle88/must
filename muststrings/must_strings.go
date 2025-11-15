// Package muststrings provides string-specific assertion utilities with panic-on-failure semantics
// Implements validation functions using Go standard strings package
// Supports length checking, prefix/suffix validation, and substring containment testing
// Integrates with zap structured logging to provide detailed context when assertions are not met
//
// muststrings 提供字符串特定的断言工具，带 panic-on-failure 语义
// 使用 Go 标准 strings 包实现字符串操作的验证函数
// 支持长度检查、前缀/后缀验证和子串包含性测试
// 与 zap 结构化日志集成，当断言不满足时提供详细上下文
package muststrings

import (
	"strings"

	"github.com/yyle88/zaplog"
	"go.uber.org/zap"
)

// Length expects the string to have length n. Panics if the length is not n.
// Length 期望字符串的长度为 n。如果长度不是 n，则触发 panic。
func Length(a string, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// Len is an abbreviation of Length, serving the same purpose. Panics if the length is not n.
// Len 是 Length 的缩写，功能相同。如果长度不是 n，则触发 panic。
func Len(a string, n int) {
	if len(a) != n {
		zaplog.ZAPS.Skip1.LOG.Panic("LENGTH MISMATCH(NOT MATCH)", zap.Int("len", len(a)), zap.Int("n", n))
	}
}

// HasPrefix checks if the string has the specified prefix, panics if not.
// HasPrefix 检查字符串是否有指定的前缀，没有则触发 panic。
func HasPrefix(a string, prefix string) {
	if !strings.HasPrefix(a, prefix) {
		zaplog.ZAPS.Skip1.LOG.Panic("STRING MISSING PREFIX(SHOULD HAVE PREFIX)", zap.String("string", a), zap.String("prefix", prefix))
	}
}

// HasSuffix checks if the string has the specified suffix, panics if not.
// HasSuffix 检查字符串是否有指定的后缀，没有则触发 panic。
func HasSuffix(a string, suffix string) {
	if !strings.HasSuffix(a, suffix) {
		zaplog.ZAPS.Skip1.LOG.Panic("STRING MISSING SUFFIX(SHOULD HAVE SUFFIX)", zap.String("string", a), zap.String("suffix", suffix))
	}
}

// NotHasPrefix checks if the string does not have the specified prefix, panics if it does.
// NotHasPrefix 检查字符串是否没有指定的前缀，有则触发 panic。
func NotHasPrefix(a string, prefix string) {
	if strings.HasPrefix(a, prefix) {
		zaplog.ZAPS.Skip1.LOG.Panic("STRING HAS PREFIX(SHOULD NOT HAVE PREFIX)", zap.String("string", a), zap.String("prefix", prefix))
	}
}

// NotHasSuffix checks if the string does not have the specified suffix, panics if it does.
// NotHasSuffix 检查字符串是否没有指定的后缀，有则触发 panic。
func NotHasSuffix(a string, suffix string) {
	if strings.HasSuffix(a, suffix) {
		zaplog.ZAPS.Skip1.LOG.Panic("STRING HAS SUFFIX(SHOULD NOT HAVE SUFFIX)", zap.String("string", a), zap.String("suffix", suffix))
	}
}

// Contains checks if the string contains the specified substring, panics if not.
// Contains 检查字符串是否包含指定的子串，没有则触发 panic。
func Contains(a string, sub string) {
	if !strings.Contains(a, sub) {
		zaplog.ZAPS.Skip1.LOG.Panic("STRING MISSING SUBSTRING(SHOULD HAVE SUBSTRING)", zap.String("string", a), zap.String("substring", sub))
	}
}

// NotContains checks if the string does not contain the specified substring, panics if it does.
// NotContains 检查字符串是否不包含指定的子串，有则触发 panic。
func NotContains(a string, sub string) {
	if strings.Contains(a, sub) {
		zaplog.ZAPS.Skip1.LOG.Panic("STRING HAS SUBSTRING(SHOULD NOT HAVE SUBSTRING)", zap.String("string", a), zap.String("substring", sub))
	}
}
