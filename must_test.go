// Package must_test provides comprehensive testing of must assertion package
// Tests include boolean assertions, error handling, value validation, and panic-on-failure semantics
// Checks each assertion function with both success and failure cases
//
// must_test 为 must 断言包提供全面的测试
// 测试涵盖布尔断言、异常处理、值验证和 panic-on-failure 语义
// 使用成功和失败案例检查每个断言函数
package must_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
)

// TestTrue tests boolean true assertion
// Validates True passes with true and panics with false
//
// TestTrue 测试布尔 true 断言
// 验证 True 在 true 时通过，在 false 时 panic
func TestTrue(t *testing.T) {
	must.True(true)

	require.Panics(t, func() {
		must.True(false)
	})
}

// TestDone tests no error assertion
// Checks Done passes with nil error and panics with non-nil error
//
// TestDone 测试无异常断言
// 检查 Done 在无异常时通过，在有异常时 panic
func TestDone(t *testing.T) {
	must.Done(error(nil))

	require.Panics(t, func() {
		must.Done(errors.New("wa"))
	})
}

// TestMust tests error absence assertion
// Checks Must passes with nil error and panics with non-nil error
//
// TestMust 测试异常缺失断言
// 检查 Must 在无异常时通过，在有异常时 panic
func TestMust(t *testing.T) {
	must.Must(error(nil))

	require.Panics(t, func() {
		must.Must(errors.New("wa"))
	})
}

// TestNice tests non-zero value assertion with return
// Validates Nice returns non-zero values and panics with zero values
//
// TestNice 测试非零值断言并返回
// 验证 Nice 返回非零值并在零值时 panic
func TestNice(t *testing.T) {
	require.Equal(t, 88, must.Nice(88))
	require.Equal(t, "xyz", must.Nice("xyz"))
	require.Equal(t, 3.1415926, must.Nice(3.1415926))

	require.Panics(t, func() {
		must.Nice(0.0)
	})
}

// TestZero tests zero value assertion
// Validates Zero passes with zero values and panics with non-zero values
//
// TestZero 测试零值断言
// 验证 Zero 在零值时通过，在非零值时 panic
func TestZero(t *testing.T) {
	must.Zero(0)
	must.Zero("")
	must.Zero(0.0)

	require.Panics(t, func() {
		must.Zero(uint64(200))
	})
}

// TestNone tests zero value assertion (alias of Zero)
// Validates None passes with zero values and panics with non-zero values
//
// TestNone 测试零值断言（Zero 的别名）
// 验证 None 在零值时通过，在非零值时 panic
func TestNone(t *testing.T) {
	must.None(0)
	must.None(0.0)
	must.None("")
	must.None(error(nil))

	require.Panics(t, func() {
		must.None(errors.New("wa"))
	})
}

// TestEquals tests value matching assertion
// Checks Equals passes when values match and panics when values mismatch
//
// TestEquals 测试值匹配断言
// 检查 Equals 在值匹配时通过，在值不同时 panic
func TestEquals(t *testing.T) {
	must.Equals("abc", "abc")
	must.Equals(123, 123)
	must.Equals(0.8, 0.8)

	require.Panics(t, func() {
		must.Equals("abc", "xyz")
	})
}

// TestSame tests value sameness assertion
// Checks Same passes when values match and panics on mismatch
//
// TestSame 测试值相同断言
// 验证 Same 在值匹配时通过，在值不同时 panic
func TestSame(t *testing.T) {
	must.Same("abc", "abc")
	must.Same(123, 123)
	must.Same(0.8, 0.8)

	require.Panics(t, func() {
		must.Same("abc", "xyz")
	})
}

// TestSameNice tests value sameness with non-zero check and return
// Checks SameNice passes when values match and non-zero, panics on mismatch
//
// TestSameNice 测试值相同且非零断言并返回
// 验证 SameNice 在值匹配且非零时通过，在不匹配或为零时 panic
func TestSameNice(t *testing.T) {
	require.Equal(t, "abc", must.SameNice("abc", "abc"))
	require.Equal(t, 123, must.SameNice(123, 123))
	require.Equal(t, 0.8, must.SameNice(0.8, 0.8))

	require.Panics(t, func() {
		must.SameNice("abc", "xyz")
	})

	require.Panics(t, func() {
		must.SameNice(0, 0)
	})
}

// TestSane tests value sameness and non-zero check with return
// Checks Sane passes when values match and non-zero, panics on mismatch
//
// TestSane 测试值相同且非零断言并返回
// 验证 Sane 在值匹配且非零时通过，在不匹配或为零时 panic
func TestSane(t *testing.T) {
	require.Equal(t, "abc", must.Sane("abc", "abc"))
	require.Equal(t, 123, must.Sane(123, 123))
	require.Equal(t, 0.8, must.Sane(0.8, 0.8))

	require.Panics(t, func() {
		must.Sane("abc", "xyz")
	})

	require.Panics(t, func() {
		must.Sane(0, 0)
	})
}

// TestDiff tests value difference assertion
// Checks Diff passes when values mismatch and panics when match
//
// TestDiff 测试值不同断言
// 验证 Diff 在值不同时通过，在值匹配时 panic
func TestDiff(t *testing.T) {
	must.Diff("abc", "xyz")
	must.Diff(123, 321)
	must.Diff(0.8, 0.88)

	require.Panics(t, func() {
		must.Diff("abc", "abc")
	})
}

// TestDifferent tests value difference assertion
// Checks Different passes when values mismatch and panics when match
//
// TestDifferent 测试值不同断言
// 验证 Different 在值不同时通过，在值匹配时 panic
func TestDifferent(t *testing.T) {
	must.Different("abc", "xyz")
	must.Different(123, 321)
	must.Different(0.8, 0.88)

	require.Panics(t, func() {
		must.Different("abc", "abc")
	})
}

// TestIs tests value identity check
// Checks Is passes when values match and panics when values mismatch
//
// TestIs 测试值一致性检查
// 检查 Is 在值匹配时通过，在值不同时 panic
func TestIs(t *testing.T) {
	must.Is(123, 123)
	must.Is("!", "!")
	must.Is('!', '!')
	must.Is(0.5, 0.5)

	require.Panics(t, func() {
		must.Is(1, 2)
	})
}

// TestIse tests error matching assertion using errors.Is
// Validates Ise passes when errors match and panics when errors differ
//
// TestIse 测试使用 errors.Is 的错误匹配断言
// 验证 Ise 在错误匹配时通过，在错误不同时 panic
func TestIse(t *testing.T) {
	errA := errors.New("a")
	must.Ise(errA, errA)

	require.Panics(t, func() {
		must.Ise(errors.New("a"), errors.New("b"))
	})
}

// TestOk tests non-zero value assertion
// Validates Ok passes with non-zero values and panics with zero values
//
// TestOk 测试非零值断言
// 验证 Ok 在非零值时通过，在零值时 panic
func TestOk(t *testing.T) {
	must.Ok(1990)

	require.Panics(t, func() {
		must.Ok("")
	})
}

// TestOK tests non-zero value assertion (uppercase variant)
// Validates OK passes with non-zero values and panics with zero values
//
// TestOK 测试非零值断言（大写变体）
// 验证 OK 在非零值时通过，在零值时 panic
func TestOK(t *testing.T) {
	must.OK("yyle")
	must.OK(88)
	must.OK(0.88)
	must.OK(uint64(88))
	must.OK('8')

	require.Panics(t, func() {
		must.OK(0)
	})
}

// TestTRUE tests boolean true assertion (uppercase variant)
// Validates TRUE passes with true and panics with false
//
// TestTRUE 测试布尔 true 断言（大写变体）
// 验证 TRUE 在 true 时通过，在 false 时 panic
func TestTRUE(t *testing.T) {
	must.TRUE(true)

	require.Panics(t, func() {
		must.TRUE(false)
	})
}

// TestFALSE tests boolean false assertion (uppercase variant)
// Validates FALSE passes with false and panics with true
//
// TestFALSE 测试布尔 false 断言（大写变体）
// 验证 FALSE 在 false 时通过，在 true 时 panic
func TestFALSE(t *testing.T) {
	must.FALSE(false)

	require.Panics(t, func() {
		must.FALSE(true)
	})
}

// TestFalse tests boolean false assertion
// Validates False passes with false and panics with true
//
// TestFalse 测试布尔 false 断言
// 验证 False 在 false 时通过，在 true 时 panic
func TestFalse(t *testing.T) {
	must.False(false)

	require.Panics(t, func() {
		must.False(true)
	})
}

// TestHave tests slice non-empty assertion
// Validates Have passes with non-empty slices and panics with empty slices
//
// TestHave 测试切片非空断言
// 验证 Have 在非空切片时通过，在空切片时 panic
func TestHave(t *testing.T) {
	must.Have([]int{0, 1, 2, 3})
	must.Have([]string{"abc", "xyz"})
	must.Have([]float64{0.0})

	require.Panics(t, func() {
		must.Have([]uint64{})
	})
}

// TestLength tests slice length assertion
// Validates Length passes when slice length matches expected value and panics on mismatch
//
// TestLength 测试切片长度断言
// 验证 Length 在切片长度匹配期望值时通过，在不匹配时 panic
func TestLength(t *testing.T) {
	must.Length([]int{0, 1, 2, 3}, 4)
	must.Length([]string{"abc", "xyz"}, 2)
	must.Length([]float64{0.0}, 1)
	must.Length([]uint64{}, 0)

	require.Panics(t, func() {
		must.Length([]any{"a", 1, 'x'}, 2)
	})
}

// TestLen tests slice length assertion (short form)
// Validates Len passes when slice length matches expected value and panics on mismatch
//
// TestLen 测试切片长度断言（简短形式）
// 验证 Len 在切片长度匹配期望值时通过，在不匹配时 panic
func TestLen(t *testing.T) {
	must.Len([]int{}, 0)
	must.Len([]float64{0.1, 0.2}, 2)

	require.Panics(t, func() {
		must.Len([]uint64{}, 88)
	})
}

// TestIn tests value membership in slice assertion
// Validates In passes when value exists in slice and panics when value not found
//
// TestIn 测试值在切片中的成员断言
// 验证 In 在值存在于切片中时通过，在未找到值时 panic
func TestIn(t *testing.T) {
	must.In(1, []int{1, 2, 3})
	must.In(2.0, []int{1, 2.0, 3})

	require.Panics(t, func() {
		must.In("a", []string{"b", "c"})
	})
}

// TestContains tests slice contains value assertion
// Validates Contains passes when slice contains value and panics when value not found
//
// TestContains 测试切片包含值断言
// 验证 Contains 在切片包含值时通过，在未找到值时 panic
func TestContains(t *testing.T) {
	must.Contains([]string{"a", "b", "c"}, "a")

	require.Panics(t, func() {
		must.Contains([]int{1, 2, 3}, 4)
	})
}

// Example test struct
// 示例测试结构体
type Example struct {
	S string // String field // 字符串字段
}

// TestNull tests nil ptr assertion
// Checks Null passes with nil ptr and panics with non-nil pts
//
// TestNull 测试 nil 指针断言
// 检查 Null 在 nil 指针时通过，在非 nil 指针时 panic
func TestNull(t *testing.T) {
	var example *Example
	must.Null(example)

	require.Panics(t, func() {
		must.Null(&Example{S: "abc"})
	})
}

// TestFull tests non-nil pts assertion with return
// Checks Full returns non-nil pts and panics with nil pts
//
// TestFull 测试非 nil 指针断言并返回
// 检查 Full 返回非 nil 指针并在 nil 指针时 panic
func TestFull(t *testing.T) {
	res := must.Full(&Example{S: "abc"})
	require.Equal(t, "abc", res.S)

	require.Panics(t, func() {
		var example *Example
		must.Full(example)
	})
}
