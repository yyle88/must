// Package utils provides internal testing of zero value utility functions
// Tests include generic zero value generation across different types
//
// utils 为零值工具函数提供内部测试
// 测试涵盖不同类型的泛型零值生成
package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TestZero tests generic zero value generation
// Validates Zero returns correct zero values of different types
//
// TestZero 测试泛型零值生成
// 验证 Zero 为不同类型返回正确的零值
func TestZero(t *testing.T) {
	require.Equal(t, 0, Zero[int]())
	require.Equal(t, "", Zero[string]())
	require.Equal(t, false, Zero[bool]())
}
