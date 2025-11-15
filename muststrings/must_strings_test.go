// Package muststrings_test provides comprehensive testing of muststrings assertion package
// Tests include string length validation, prefix and suffix checks, and substring containment
// Checks each assertion functions with both success and failure cases
//
// muststrings_test 为 muststrings 断言包提供全面的测试
// 测试涵盖字符串长度验证、前缀和后缀检查以及子串包含性
// 使用成功和失败案例验证所有断言函数
package muststrings_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/muststrings"
)

// TestLength tests string length assertion
// Validates Length passes when string length matches expected value and panics on mismatch
//
// TestLength 测试字符串长度断言
// 验证 Length 在字符串长度匹配期望值时通过，在不匹配时 panic
func TestLength(t *testing.T) {
	muststrings.Length("abc", 3)

	require.Panics(t, func() {
		muststrings.Length("xyz", 0)
	})
}

// TestLen tests string length assertion (short form)
// Validates Len passes when string length matches expected value and panics on mismatch
//
// TestLen 测试字符串长度断言（简短形式）
// 验证 Len 在字符串长度匹配期望值时通过，在不匹配时 panic
func TestLen(t *testing.T) {
	muststrings.Len("123", 3)

	require.Panics(t, func() {
		muststrings.Len("000", 0)
	})
}

// TestHasPrefix tests string prefix presence assertion
// Validates HasPrefix passes when string has prefix and panics when prefix not found
//
// TestHasPrefix 测试字符串前缀存在断言
// 验证 HasPrefix 在字符串有前缀时通过，在前缀不存在时 panic
func TestHasPrefix(t *testing.T) {
	muststrings.HasPrefix("hello", "he")

	require.Panics(t, func() {
		muststrings.HasPrefix("hello", "hi")
	})
}

// TestHasSuffix tests string suffix presence assertion
// Validates HasSuffix passes when string has suffix and panics when suffix not found
//
// TestHasSuffix 测试字符串后缀存在断言
// 验证 HasSuffix 在字符串有后缀时通过，在后缀不存在时 panic
func TestHasSuffix(t *testing.T) {
	muststrings.HasSuffix("world", "ld")

	require.Panics(t, func() {
		muststrings.HasSuffix("world", "lo")
	})
}

// TestNotHasPrefix tests string prefix absence assertion
// Validates NotHasPrefix passes when string lacks prefix and panics when prefix found
//
// TestNotHasPrefix 测试字符串前缀不存在断言
// 验证 NotHasPrefix 在字符串无前缀时通过，在前缀存在时 panic
func TestNotHasPrefix(t *testing.T) {
	muststrings.NotHasPrefix("hello", "hi")

	require.Panics(t, func() {
		muststrings.NotHasPrefix("hello", "he")
	})
}

// TestNotHasSuffix tests string suffix absence assertion
// Validates NotHasSuffix passes when string lacks suffix and panics when suffix found
//
// TestNotHasSuffix 测试字符串后缀不存在断言
// 验证 NotHasSuffix 在字符串无后缀时通过，在后缀存在时 panic
func TestNotHasSuffix(t *testing.T) {
	muststrings.NotHasSuffix("world", "lo")

	require.Panics(t, func() {
		muststrings.NotHasSuffix("world", "ld")
	})
}

// TestContains tests string substring containment assertion
// Validates Contains passes when string contains substring and panics when substring not found
//
// TestContains 测试字符串子串包含断言
// 验证 Contains 在字符串包含子串时通过，在子串不存在时 panic
func TestContains(t *testing.T) {
	muststrings.Contains("hello world", "world")

	require.Panics(t, func() {
		muststrings.Contains("hello world", "planet")
	})
}

// TestNotContains tests string substring absence assertion
// Validates NotContains passes when string lacks substring and panics when substring found
//
// TestNotContains 测试字符串子串不包含断言
// 验证 NotContains 在字符串不包含子串时通过，在子串存在时 panic
func TestNotContains(t *testing.T) {
	muststrings.NotContains("hello world", "planet")

	require.Panics(t, func() {
		muststrings.NotContains("hello world", "world")
	})
}
