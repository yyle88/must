// Package utils provides support functions enabling the must assertion packages
// Implements generic support functions used across different assertion modules
// Not intended to be used outside this module, serves as support mechanism
//
// utils 提供内部工具函数，支持 must 断言包
// 实现在不同断言模块间使用的泛型辅助函数
// 不在此模块外使用，作为支持机制
package utils

// Zero returns the zero value of type T using named return value initialization.
// Zero 使用命名返回值初始化返回类型 T 的零值。
func Zero[T any]() (x T) {
	return x
}
