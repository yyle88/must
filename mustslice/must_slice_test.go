package mustslice_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/mustslice"
)

func TestEquals(t *testing.T) {
	// 正常情况下，两个切片相等
	mustslice.Equals([]int{1, 2, 3}, []int{1, 2, 3})

	// 切片不相等时触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.Equals([]string{"a"}, []string{"b"})
	})
}

func TestDifferent(t *testing.T) {
	// 正常情况下，两个切片相等
	mustslice.Different([]string{"a"}, []string{"b"})

	// 切片不相等时触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.Different([]int{1, 2, 3}, []int{1, 2, 3})
	})
}

func TestContains(t *testing.T) {
	// 切片包含指定元素时通过
	mustslice.Contains([]string{"a", "b", "c"}, "a")

	// 切片不包含指定元素时触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.Contains([]int{1, 2, 3}, 4)
	})
}

func TestIn(t *testing.T) {
	// 元素在切片中时通过
	mustslice.In("a", []string{"a", "b", "c"})

	// 元素不在切片中时触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.In(4, []int{1, 2, 3})
	})
}

func TestNice(t *testing.T) {
	// 非空切片返回自身
	require.Equal(t, []int{1, 2, 3}, mustslice.Nice([]int{1, 2, 3}))

	// 空切片时触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.Nice([]string{})
	})
}

func TestNone(t *testing.T) {
	mustslice.None([]int{})

	tests.ExpectPanic(t, func() {
		mustslice.None([]int{1, 2, 3})
	})
}

func TestHave(t *testing.T) {
	// 非空切片通过
	mustslice.Have([]int{1, 2, 3})

	// 空切片触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.Have([]string{})
	})
}

func TestLength(t *testing.T) {
	// 切片长度符合期望值时通过
	mustslice.Length([]int{1, 2, 3}, 3)

	// 切片长度不符合期望值时触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.Length([]string{"a", "b"}, 3)
	})
}

func TestLen(t *testing.T) {
	// Len 与 Length 作用相同，测试该函数的行为
	mustslice.Len([]int{1, 2, 3}, 3)

	// 切片长度不符合期望值时触发 panic
	tests.ExpectPanic(t, func() {
		mustslice.Len([]string{"a", "b"}, 3)
	})
}
