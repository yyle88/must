// Package mustnum_test provides comprehensive testing of mustnum numeric assertion package
// Tests include numeric comparisons, range validations, and sign checks
// Checks each assertion functions with both success and failure cases
//
// mustnum_test 为 mustnum 数值断言包提供全面的测试
// 测试涵盖数值比较、范围验证和符号检查
// 使用成功和失败案例验证所有断言函数
package mustnum_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/mustnum"
)

// TestLess tests numeric less-than assertion
// Validates Less passes when first value is less than second and panics when not
//
// TestLess 测试数值小于断言
// 验证 Less 在第一个值小于第二个值时通过，否则 panic
func TestLess(t *testing.T) {
	mustnum.Less(100, 200)
	mustnum.Less(0.1, 0.2)

	require.Panics(t, func() {
		mustnum.Less(2, 1)
	})

	require.Panics(t, func() {
		mustnum.Less(0.90, 0.11)
	})
}

// TestLt tests numeric less-than assertion (short form)
// Validates Lt passes when first value is less than second and panics when not
//
// TestLt 测试数值小于断言（简短形式）
// 验证 Lt 在第一个值小于第二个值时通过，否则 panic
func TestLt(t *testing.T) {
	mustnum.Lt(100, 200)
	mustnum.Lt(0.1, 0.2)

	require.Panics(t, func() {
		mustnum.Lt(2, 1)
	})

	require.Panics(t, func() {
		mustnum.Lt(0.90, 0.11)
	})
}

// TestLte tests numeric less-than-or-at-most assertion
// Validates Lte passes when first value is less than / at most second and panics when greater
//
// TestLte 测试数值小于或等于断言
// 验证 Lte 在第一个值小于或等于第二个值时通过，在大于时 panic
func TestLte(t *testing.T) {
	mustnum.Lte(100, 200)
	mustnum.Lte(5, 5)
	mustnum.Lte(0.1, 0.2)

	require.Panics(t, func() {
		mustnum.Lte(2, 1)
	})

	require.Panics(t, func() {
		mustnum.Lte(0.90, 0.11)
	})
}

// TestGt tests numeric greater-than assertion
// Validates Gt passes when first value exceeds second and panics when not
//
// TestGt 测试数值大于断言
// 验证 Gt 在第一个值大于第二个值时通过，否则 panic
func TestGt(t *testing.T) {
	mustnum.Gt(200, 100)
	mustnum.Gt(0.2, 0.1)

	require.Panics(t, func() {
		mustnum.Gt(1, 2)
	})

	require.Panics(t, func() {
		mustnum.Gt(0.11, 0.90)
	})
}

// TestGte tests numeric greater-than-or-at-least assertion
// Validates Gte passes when first value exceeds / matches second and panics when less
//
// TestGte 测试数值大于或等于断言
// 验证 Gte 在第一个值大于或等于第二个值时通过，在小于时 panic
func TestGte(t *testing.T) {
	mustnum.Gte(200, 100)
	mustnum.Gte(5, 5)
	mustnum.Gte(0.2, 0.1)

	require.Panics(t, func() {
		mustnum.Gte(1, 2)
	})

	require.Panics(t, func() {
		mustnum.Gte(0.11, 0.90)
	})
}

// TestNice tests non-zero numeric value assertion
// Validates Nice passes with non-zero values and panics with zero
//
// TestNice 测试非零数值断言
// 验证 Nice 在非零值时通过，在零值时 panic
func TestNice(t *testing.T) {
	mustnum.Nice(100)
	mustnum.Nice(-100)
	mustnum.Nice(0.1)
	mustnum.Nice(-0.1)

	require.Panics(t, func() {
		mustnum.Nice(0)
	})

	require.Panics(t, func() {
		mustnum.Nice(0.0)
	})
}

// TestZero tests zero numeric value assertion
// Validates Zero passes with zero values and panics with non-zero values
//
// TestZero 测试零值数值断言
// 验证 Zero 在零值时通过，在非零值时 panic
func TestZero(t *testing.T) {
	mustnum.Zero(0)
	mustnum.Zero(0.0)

	require.Panics(t, func() {
		mustnum.Zero(100)
	})

	require.Panics(t, func() {
		mustnum.Zero(-100)
	})

	require.Panics(t, func() {
		mustnum.Zero(0.1)
	})

	require.Panics(t, func() {
		mustnum.Zero(-0.1)
	})
}

// TestPositive tests positive numeric value assertion
// Validates Positive passes when value exceeds zero and panics when value is zero / negative
//
// TestPositive 测试正数值断言
// 验证 Positive 在值大于零时通过，在值为零或负数时 panic
func TestPositive(t *testing.T) {
	mustnum.Positive(100)
	mustnum.Positive(0.1)
	mustnum.Positive(uint64(1))

	require.Panics(t, func() {
		mustnum.Positive(0)
	})

	require.Panics(t, func() {
		mustnum.Positive(-1)
	})

	require.Panics(t, func() {
		mustnum.Positive(-0.1)
	})
}

// TestNegative tests negative numeric value assertion
// Validates Negative passes when value is below zero and panics when value is zero / positive
//
// TestNegative 测试负数值断言
// 验证 Negative 在值小于零时通过，在值为零或正数时 panic
func TestNegative(t *testing.T) {
	mustnum.Negative(-100)
	mustnum.Negative(-0.1)
	mustnum.Negative(int64(-1))

	require.Panics(t, func() {
		mustnum.Negative(0)
	})

	require.Panics(t, func() {
		mustnum.Negative(1)
	})

	require.Panics(t, func() {
		mustnum.Negative(0.1)
	})
}
