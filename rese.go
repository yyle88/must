package must

func V0(err error) {
	Must(err)
}

func V1[T1 any](v1 T1, err error) T1 {
	Must(err)
	return v1
}

func V2[T1, T2 any](v1 T1, v2 T2, err error) (T1, T2) {
	Must(err)
	return v1, v2
}

func P0(err error) {
	Must(err)
}

func P1[T1 any](v1 *T1, err error) *T1 {
	Must(err)
	Full(v1)
	return v1
}

func P2[T1, T2 any](v1 *T1, v2 *T2, err error) (*T1, *T2) {
	Must(err)
	Full(v1)
	Full(v2)
	return v1, v2
}

func C0(err error) {
	Must(err)
}

func C1[T1 comparable](v1 T1, err error) T1 {
	Must(err)
	Nice(v1)
	return v1
}

func C2[T1, T2 comparable](v1 T1, v2 T2, err error) (T1, T2) {
	Must(err)
	Nice(v1)
	Nice(v2)
	return v1, v2
}
