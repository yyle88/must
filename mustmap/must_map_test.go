package mustmap_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/mustmap"
)

func TestEquals(t *testing.T) {
	mustmap.Equals(map[int]string{
		1: "a",
		2: "b",
	}, map[int]string{
		2: "b",
		1: "a",
	})

	tests.ExpectPanic(t, func() {
		mustmap.Equals(map[string]int{
			"a": 1,
		}, map[string]int{
			"b": 2,
		})
	})
}

func TestDiff(t *testing.T) {
	mustmap.Diff(map[string]int{
		"a": 1,
	}, map[string]int{
		"b": 2,
	})

	tests.ExpectPanic(t, func() {
		mustmap.Diff(map[int]string{
			1: "a",
			2: "b",
		}, map[int]string{
			2: "b",
			1: "a",
		})
	})
}

func TestDifferent(t *testing.T) {
	mustmap.Different(map[string]int{
		"a": 1,
	}, map[string]int{
		"b": 2,
	})

	tests.ExpectPanic(t, func() {
		mustmap.Different(map[int]string{
			1: "a",
			2: "b",
		}, map[int]string{
			2: "b",
			1: "a",
		})
	})
}

func TestHave(t *testing.T) {
	// 正常情况: 非空 map 不应触发 panic
	mustmap.Have(map[int]string{
		1: "value1",
		2: "value2",
	})

	// 异常情况: 空 map 应触发 panic
	tests.ExpectPanic(t, func() {
		mustmap.Have(map[string]int{})
	})
}

func TestNice(t *testing.T) {
	require.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, mustmap.Nice(map[string]int{"c": 3, "b": 2, "a": 1}))

	tests.ExpectPanic(t, func() {
		mustmap.Nice(map[string]int{})
	})
}

func TestZero(t *testing.T) {
	mustmap.Zero(map[int]string{})
	mustmap.Zero(map[string]int{})
	mustmap.Zero((map[int]string)(nil))

	tests.ExpectPanic(t, func() {
		mustmap.Zero(map[int]string{1: "a"})
	})

	tests.ExpectPanic(t, func() {
		mustmap.Zero(map[string]int{"a": 1})
	})

	tests.ExpectPanic(t, func() {
		mustmap.Zero(map[string]int{"a": 1, "b": 2})
	})
}

func TestNone(t *testing.T) {
	mustmap.None(map[int]int{})

	tests.ExpectPanic(t, func() {
		mustmap.None(map[string]int{"a": 1, "b": 2})
	})
}

func TestLength(t *testing.T) {
	// 正常情况: map 的长度等于期望值
	mustmap.Length(map[int]string{
		1: "value1",
		2: "value2",
	}, 2)

	// 异常情况: map 的长度不等于期望值
	tests.ExpectPanic(t, func() {
		mustmap.Length(map[int]string{
			1: "value1",
		}, 2)
	})

	// 边界情况: 空 map 长度为 0
	mustmap.Length(map[string]int{}, 0)

	// 异常情况: 非空 map 长度与期望值不符
	tests.ExpectPanic(t, func() {
		mustmap.Length(map[string]int{"a": 1, "b": 2}, 3)
	})
}

func TestLen(t *testing.T) {
	// 正常情况: Len 的行为与 Length 相同
	mustmap.Len(map[int]string{
		1: "value1",
		2: "value2",
	}, 2)

	// 异常情况: 长度不符触发 panic
	tests.ExpectPanic(t, func() {
		mustmap.Len(map[int]string{
			1: "value1",
		}, 2)
	})

	// 边界情况: 空 map 长度为 0
	mustmap.Len(map[string]int{}, 0)

	// 异常情况: 非空 map 长度与期望值不符
	tests.ExpectPanic(t, func() {
		mustmap.Len(map[string]int{"a": 1, "b": 2}, 3)
	})
}

func TestGet(t *testing.T) {
	// 测试键存在时，返回对应的值
	value := mustmap.Get(map[string]int{"a": 1, "b": 2}, "a")
	require.Equal(t, 1, value)

	// 测试键不存在时，触发 panic
	tests.ExpectPanic(t, func() {
		mustmap.Get(map[string]int{"a": 1, "b": 2}, "c")
	})
}
