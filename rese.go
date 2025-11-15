package must

// V0 validates no error occurred. Panics if error is non-nil.
// V0 验证没有错误发生。如果错误非 nil 则触发 panic。
func V0(err error) {
	Must(err)
}

// V1 validates no error and returns the value. Panics if error is non-nil.
// V1 验证没有错误并返回值。如果错误非 nil 则触发 panic。
func V1[T1 any](v1 T1, err error) T1 {
	Must(err)
	return v1
}

// V2 validates no error and returns two values. Panics if error is non-nil.
// V2 验证没有错误并返回两个值。如果错误非 nil 则触发 panic。
func V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	Must(err)
	return v1, v2
}

// P0 validates no error occurred. Panics if error is non-nil.
// P0 验证没有错误发生。如果错误非 nil 则触发 panic。
func P0(err error) {
	Must(err)
}

// P1 validates no error and returns non-nil data. Panics if error is non-nil / data is nil.
// P1 验证没有错误并返回非 nil 指针。如果错误非 nil 或指针为 nil 则触发 panic。
func P1[T1 any](v1 *T1, err error) *T1 {
	Must(err)
	Full(v1)
	return v1
}

// P2 validates no error and returns two non-nil data. Panics if error is non-nil / any data is nil.
// P2 验证没有错误并返回两个非 nil 指针。如果错误非 nil 或任何指针为 nil 则触发 panic。
func P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2) {
	Must(err)
	Full(v1)
	Full(v2)
	return v1, v2
}

// C0 validates no error occurred. Panics if error is non-nil.
// C0 验证没有错误发生。如果错误非 nil 则触发 panic。
func C0(err error) {
	Must(err)
}

// C1 validates no error and returns non-zero comparable value. Panics if error is non-nil / value is zero.
// C1 验证没有错误并返回非零可比较值。如果错误非 nil 或值为零则触发 panic。
func C1[T1 comparable](v1 T1, err error) T1 {
	Must(err)
	Nice(v1)
	return v1
}

// C2 validates no error and returns two non-zero comparable values. Panics if error is non-nil / any value is zero.
// C2 验证没有错误并返回两个非零可比较值。如果错误非 nil 或任何值为零则触发 panic。
func C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2) {
	Must(err)
	Nice(v1)
	Nice(v2)
	return v1, v2
}
