// Package mustslice_test provides comprehensive testing of mustslice assertion package
// Tests include slice equality, difference validation, element containment, and length checking
// Checks each assertion functions with both success and failure cases
//
// mustslice_test 为 mustslice 断言包提供全面的测试
// 测试涵盖切片相等性、差异验证、元素包含性和长度检查
// 使用成功和失败案例验证所有断言函数
package mustslice_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/mustslice"
)

// TestEquals tests slice equality assertion
// Validates Equals passes when slices have same content and panics when slices differ
//
// TestEquals 测试切片相等断言
// 验证 Equals 在切片内容相同时通过，在切片不同时 panic
func TestEquals(t *testing.T) {
	mustslice.Equals([]int{1, 2, 3}, []int{1, 2, 3})

	require.Panics(t, func() {
		mustslice.Equals([]string{"a"}, []string{"b"})
	})
}

// TestDiff tests slice difference assertion
// Validates Diff passes when slices differ and panics when slices are same
//
// TestDiff 测试切片差异断言
// 验证 Diff 在切片不同时通过，在切片相同时 panic
func TestDiff(t *testing.T) {
	mustslice.Diff([]string{"a"}, []string{"b"})

	require.Panics(t, func() {
		mustslice.Diff([]int{1, 2, 3}, []int{1, 2, 3})
	})
}

// TestDifferent tests slice difference assertion
// Validates Different passes when slices differ and panics when slices are same
//
// TestDifferent 测试切片差异断言
// 验证 Different 在切片不同时通过，在切片相同时 panic
func TestDifferent(t *testing.T) {
	mustslice.Different([]string{"a"}, []string{"b"})

	require.Panics(t, func() {
		mustslice.Different([]int{1, 2, 3}, []int{1, 2, 3})
	})
}

// TestContains tests slice contains element assertion
// Validates Contains passes when slice contains element and panics when element not found
//
// TestContains 测试切片包含元素断言
// 验证 Contains 在切片包含元素时通过，在未找到元素时 panic
func TestContains(t *testing.T) {
	mustslice.Contains([]string{"a", "b", "c"}, "a")

	require.Panics(t, func() {
		mustslice.Contains([]int{1, 2, 3}, 4)
	})
}

// TestIn tests element membership in slice assertion
// Validates In passes when element exists in slice and panics when element not found
//
// TestIn 测试元素在切片中的成员断言
// 验证 In 在元素存在于切片中时通过，在未找到元素时 panic
func TestIn(t *testing.T) {
	mustslice.In("a", []string{"a", "b", "c"})

	require.Panics(t, func() {
		mustslice.In(4, []int{1, 2, 3})
	})
}

// TestNice tests non-empty slice assertion with return
// Validates Nice returns non-empty slices and panics with empty slices
//
// TestNice 测试非空切片断言并返回
// 验证 Nice 返回非空切片并在空切片时 panic
func TestNice(t *testing.T) {
	require.Equal(t, []int{1, 2, 3}, mustslice.Nice([]int{1, 2, 3}))

	require.Panics(t, func() {
		mustslice.Nice([]string{})
	})
}

// TestZero tests empty slice assertion
// Validates Zero passes with empty slices and panics with non-empty slices
//
// TestZero 测试空切片断言
// 验证 Zero 在空切片时通过，在非空切片时 panic
func TestZero(t *testing.T) {
	mustslice.Zero([]int{})
	mustslice.Zero([]string{})
	mustslice.Zero(([]int)(nil))

	require.Panics(t, func() {
		mustslice.Zero([]int{1})
	})

	require.Panics(t, func() {
		mustslice.Zero([]string{"a"})
	})

	require.Panics(t, func() {
		mustslice.Zero([]int{1, 2, 3})
	})
}

// TestNone tests empty slice assertion (alias of Zero)
// Validates None passes with empty slices and panics with non-empty slices
//
// TestNone 测试空切片断言（Zero 的别名）
// 验证 None 在空切片时通过，在非空切片时 panic
func TestNone(t *testing.T) {
	mustslice.None([]int{})

	require.Panics(t, func() {
		mustslice.None([]int{1, 2, 3})
	})
}

// TestHave tests slice non-empty assertion
// Validates Have passes with non-empty slices and panics with empty slices
//
// TestHave 测试切片非空断言
// 验证 Have 在非空切片时通过，在空切片时 panic
func TestHave(t *testing.T) {
	mustslice.Have([]int{1, 2, 3})

	require.Panics(t, func() {
		mustslice.Have([]string{})
	})
}

// TestLength tests slice length assertion
// Validates Length passes when slice length matches expected value and panics on mismatch
//
// TestLength 测试切片长度断言
// 验证 Length 在切片长度匹配期望值时通过，在不匹配时 panic
func TestLength(t *testing.T) {
	mustslice.Length([]int{1, 2, 3}, 3)

	require.Panics(t, func() {
		mustslice.Length([]string{"a", "b"}, 3)
	})
}

// TestLen tests slice length assertion (short form)
// Validates Len passes when slice length matches expected value and panics on mismatch
//
// TestLen 测试切片长度断言（简短形式）
// 验证 Len 在切片长度匹配期望值时通过，在不匹配时 panic
func TestLen(t *testing.T) {
	mustslice.Len([]int{1, 2, 3}, 3)

	require.Panics(t, func() {
		mustslice.Len([]string{"a", "b"}, 3)
	})
}
