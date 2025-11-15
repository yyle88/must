// Package mustsecret_test provides comprehensive testing of mustsecret assertion package
// Tests include secret value validation without logging sensitive data
// Checks each assertion functions with both success and failure cases
//
// mustsecret_test 为 mustsecret 断言包提供全面的测试
// 测试涵盖秘密值验证，不记录敏感数据
// 使用成功和失败案例验证所有断言函数
package mustsecret_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/mustsecret"
)

// TestNice tests non-zero secret value assertion
// Validates Nice returns non-zero values and panics with zero values without logging data
//
// TestNice 测试非零秘密值断言
// 验证 Nice 返回非零值并在零值时 panic，不记录数据
func TestNice(t *testing.T) {
	require.Equal(t, int8(5), mustsecret.Nice(int8(5)))
	require.Equal(t, "hello", mustsecret.Nice("hello"))
	require.Equal(t, 1.23, mustsecret.Nice(1.23))

	require.Panics(t, func() {
		mustsecret.Nice("")
	})
}

// TestZero tests zero secret value assertion
// Validates Zero passes with zero values and panics with non-zero values without logging data
//
// TestZero 测试零秘密值断言
// 验证 Zero 在零值时通过，在非零值时 panic，不记录数据
func TestZero(t *testing.T) {
	mustsecret.Zero(uint32(0))
	mustsecret.Zero(float32(0))
	mustsecret.Zero("")

	require.Panics(t, func() {
		mustsecret.Zero("some-value")
	})
}

// TestSame_NewData tests secret value sameness assertion
// Validates Same passes when values match and panics when values differ without logging data
//
// TestSame_NewData 测试秘密值相同断言
// 验证 Same 在值匹配时通过，在值不同时 panic，不记录数据
func TestSame_NewData(t *testing.T) {
	mustsecret.Same(true, true)
	mustsecret.Same(uint8(255), uint8(255))
	mustsecret.Same("hello", "hello")
	mustsecret.Same(float32(3.14), float32(3.14))

	require.Panics(t, func() {
		mustsecret.Same(100, 200)
	})

	require.Panics(t, func() {
		mustsecret.Same(false, true)
	})
}

// TestSane tests secret value sameness with non-zero check and return
// Validates Sane passes when values match and non-zero, panics on mismatch without logging data
//
// TestSane 测试秘密值相同且非零断言并返回
// 验证 Sane 在值匹配且非零时通过，在不匹配或为零时 panic，不记录数据
func TestSane(t *testing.T) {
	require.Equal(t, "abc", mustsecret.Sane("abc", "abc"))
	require.Equal(t, 123, mustsecret.Sane(123, 123))
	require.Equal(t, 0.8, mustsecret.Sane(0.8, 0.8))

	require.Panics(t, func() {
		mustsecret.Sane("abc", "xyz")
	})

	require.Panics(t, func() {
		mustsecret.Sane(0, 0)
	})
}
