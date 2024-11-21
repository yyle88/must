package must_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must"
)

func TestV0(t *testing.T) {
	run := func() error {
		return nil
	}

	must.V0(run())
}

func TestV1(t *testing.T) {
	run := func() (string, error) {
		return "a", nil
	}

	v1 := must.V1(run())
	require.Equal(t, "a", v1)
}

func TestV2(t *testing.T) {
	run := func() (string, uint64, error) {
		return "a", 2, nil
	}

	v1, v2 := must.V2(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
}

func TestV3(t *testing.T) {
	run := func() (string, uint64, rune, error) {
		return "a", 2, 'x', nil
	}

	v1, v2, v3 := must.V3(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
}

func TestV4(t *testing.T) {
	run := func() (string, uint64, rune, int32, error) {
		return "a", 2, 'x', 42, nil
	}

	v1, v2, v3, v4 := must.V4(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
}

func TestV5(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, error) {
		return "a", 2, 'x', 42, 3.14, nil
	}

	v1, v2, v3, v4, v5 := must.V5(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
}

func TestV6(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, error) {
		return "a", 2, 'x', 42, 3.14, true, nil
	}

	v1, v2, v3, v4, v5, v6 := must.V6(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
}

func TestV7(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, byte, error) {
		return "a", 2, 'x', 42, 3.14, true, byte(255), nil
	}

	v1, v2, v3, v4, v5, v6, v7 := must.V7(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
	require.Equal(t, byte(255), v7)
}

func TestV8(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, byte, int64, error) {
		return "a", 2, 'x', 42, 3.14, true, byte(255), int64(123456789), nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := must.V8(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
	require.Equal(t, byte(255), v7)
	require.Equal(t, int64(123456789), v8)
}

func TestV9(t *testing.T) {
	run := func() (string, uint64, rune, int32, float64, bool, byte, int64, uint32, error) {
		return "a", 2, 'x', 42, 3.14, true, byte(255), int64(123456789), uint32(987654321), nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := must.V9(run())
	require.Equal(t, "a", v1)
	require.Equal(t, uint64(2), v2)
	require.Equal(t, 'x', v3)
	require.Equal(t, int32(42), v4)
	require.Equal(t, 3.14, v5)
	require.Equal(t, true, v6)
	require.Equal(t, byte(255), v7)
	require.Equal(t, int64(123456789), v8)
	require.Equal(t, uint32(987654321), v9)
}

func TestP0(t *testing.T) {
	run := func() error {
		return nil
	}

	must.P0(run())
}

func TestP1(t *testing.T) {
	run := func() (*string, error) {
		v1 := "hello"
		return &v1, nil
	}

	p1 := must.P1(run())
	require.Equal(t, "hello", *p1)
}

func TestP2(t *testing.T) {
	run := func() (*int, *float64, error) {
		v1 := 42
		v2 := 3.14
		return &v1, &v2, nil
	}

	v1, v2 := must.P2(run())
	require.Equal(t, 42, *v1)
	require.Equal(t, 3.14, *v2)
}

func TestP3(t *testing.T) {
	run := func() (*string, *int, *bool, error) {
		v1 := "test"
		v2 := 100
		v3 := true
		return &v1, &v2, &v3, nil
	}

	v1, v2, v3 := must.P3(run())
	require.Equal(t, "test", *v1)
	require.Equal(t, 100, *v2)
	require.Equal(t, true, *v3)
}

func TestP4(t *testing.T) {
	run := func() (*int, *float64, *bool, *string, error) {
		v1 := 42
		v2 := 3.14
		v3 := false
		v4 := "world"
		return &v1, &v2, &v3, &v4, nil
	}

	v1, v2, v3, v4 := must.P4(run())
	require.Equal(t, 42, *v1)
	require.Equal(t, 3.14, *v2)
	require.Equal(t, false, *v3)
	require.Equal(t, "world", *v4)
}

func TestP5(t *testing.T) {
	run := func() (*string, *int, *bool, *float64, *rune, error) {
		v1 := "test"
		v2 := 123
		v3 := true
		v4 := 2.718
		v5 := 'a'
		return &v1, &v2, &v3, &v4, &v5, nil
	}

	v1, v2, v3, v4, v5 := must.P5(run())
	require.Equal(t, "test", *v1)
	require.Equal(t, 123, *v2)
	require.Equal(t, true, *v3)
	require.Equal(t, 2.718, *v4)
	require.Equal(t, 'a', *v5)
}

func TestP6(t *testing.T) {
	run := func() (*float64, *int, *string, *bool, *rune, *int64, error) {
		v1 := 3.14
		v2 := 42
		v3 := "hello"
		v4 := true
		v5 := 'x'
		v6 := int64(100000)
		return &v1, &v2, &v3, &v4, &v5, &v6, nil
	}

	v1, v2, v3, v4, v5, v6 := must.P6(run())
	require.Equal(t, 3.14, *v1)
	require.Equal(t, 42, *v2)
	require.Equal(t, "hello", *v3)
	require.Equal(t, true, *v4)
	require.Equal(t, 'x', *v5)
	require.Equal(t, int64(100000), *v6)
}

func TestP7(t *testing.T) {
	run := func() (*bool, *string, *int, *float64, *rune, *int64, *uint32, error) {
		v1 := true
		v2 := "world"
		v3 := 50
		v4 := 1.618
		v5 := 'y'
		v6 := int64(200000)
		v7 := uint32(300)
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, nil
	}

	v1, v2, v3, v4, v5, v6, v7 := must.P7(run())
	require.Equal(t, true, *v1)
	require.Equal(t, "world", *v2)
	require.Equal(t, 50, *v3)
	require.Equal(t, 1.618, *v4)
	require.Equal(t, 'y', *v5)
	require.Equal(t, int64(200000), *v6)
	require.Equal(t, uint32(300), *v7)
}

func TestP8(t *testing.T) {
	run := func() (*string, *int, *bool, *float64, *rune, *int64, *uint32, *uint8, error) {
		v1 := "abc"
		v2 := 999
		v3 := false
		v4 := 0.577
		v5 := 'z'
		v6 := int64(500000)
		v7 := uint32(400)
		v8 := uint8(255)
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := must.P8(run())
	require.Equal(t, "abc", *v1)
	require.Equal(t, 999, *v2)
	require.Equal(t, false, *v3)
	require.Equal(t, 0.577, *v4)
	require.Equal(t, 'z', *v5)
	require.Equal(t, int64(500000), *v6)
	require.Equal(t, uint32(400), *v7)
	require.Equal(t, uint8(255), *v8)
}

func TestP9(t *testing.T) {
	run := func() (*string, *int, *bool, *float64, *rune, *int64, *uint32, *uint8, *uint64, error) {
		v1 := "final"
		v2 := 888
		v3 := true
		v4 := 2.718
		v5 := 'p'
		v6 := int64(900000)
		v7 := uint32(500)
		v8 := uint8(100)
		v9 := uint64(1000000)
		return &v1, &v2, &v3, &v4, &v5, &v6, &v7, &v8, &v9, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := must.P9(run())
	require.Equal(t, "final", *v1)
	require.Equal(t, 888, *v2)
	require.Equal(t, true, *v3)
	require.Equal(t, 2.718, *v4)
	require.Equal(t, 'p', *v5)
	require.Equal(t, int64(900000), *v6)
	require.Equal(t, uint32(500), *v7)
	require.Equal(t, uint8(100), *v8)
	require.Equal(t, uint64(1000000), *v9)
}

func TestC0(t *testing.T) {
	run := func() error {
		return nil
	}

	must.C0(run())
}

func TestC1(t *testing.T) {
	run := func() (string, error) {
		return "hello", nil
	}

	v1 := must.C1(run())
	require.Equal(t, "hello", v1)
}

func TestC2(t *testing.T) {
	run := func() (int, float64, error) {
		return 42, 3.14, nil
	}

	v1, v2 := must.C2(run())
	require.Equal(t, 42, v1)
	require.Equal(t, 3.14, v2)
}

func TestC3(t *testing.T) {
	run := func() (bool, byte, rune, error) {
		return true, 'A', '中', nil
	}

	v1, v2, v3 := must.C3(run())
	require.True(t, v1)
	require.Equal(t, byte('A'), v2)
	require.Equal(t, '中', v3)
}

func TestC4(t *testing.T) {
	run := func() (string, int, uint, float32, error) {
		return "world", -7, 7, 1.23, nil
	}

	v1, v2, v3, v4 := must.C4(run())
	require.Equal(t, "world", v1)
	require.Equal(t, -7, v2)
	require.Equal(t, uint(7), v3)
	require.Equal(t, float32(1.23), v4)
}

func TestC5(t *testing.T) {
	run := func() (uint8, uint16, uint32, uint64, string, error) {
		return 8, 16, 32, 64, "test", nil
	}

	v1, v2, v3, v4, v5 := must.C5(run())
	require.Equal(t, uint8(8), v1)
	require.Equal(t, uint16(16), v2)
	require.Equal(t, uint32(32), v3)
	require.Equal(t, uint64(64), v4)
	require.Equal(t, "test", v5)
}

func TestC6(t *testing.T) {
	run := func() (float64, complex64, complex128, bool, string, int, error) {
		return 6.28, complex(1, 2), complex(3, 4), true, "check", -1, nil
	}

	v1, v2, v3, v4, v5, v6 := must.C6(run())
	require.Equal(t, 6.28, v1)
	require.Equal(t, complex64(complex(1, 2)), v2)
	require.Equal(t, complex(3, 4), v3)
	require.True(t, v4)
	require.Equal(t, "check", v5)
	require.Equal(t, -1, v6)
}

func TestC7(t *testing.T) {
	run := func() (uint, uint8, uint16, uint32, uint64, int, int8, error) {
		return 1, 2, 3, 4, 5, 6, 7, nil
	}

	v1, v2, v3, v4, v5, v6, v7 := must.C7(run())
	require.Equal(t, uint(1), v1)
	require.Equal(t, uint8(2), v2)
	require.Equal(t, uint16(3), v3)
	require.Equal(t, uint32(4), v4)
	require.Equal(t, uint64(5), v5)
	require.Equal(t, 6, v6)
	require.Equal(t, int8(7), v7)
}

func TestC8(t *testing.T) {
	run := func() (string, rune, byte, bool, float32, float64, int32, int64, error) {
		return "go", '中', 'B', true, 2.71, 1.618, -32, -64, nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8 := must.C8(run())
	require.Equal(t, "go", v1)
	require.Equal(t, '中', v2)
	require.Equal(t, byte('B'), v3)
	require.True(t, v4)
	require.Equal(t, float32(2.71), v5)
	require.Equal(t, 1.618, v6)
	require.Equal(t, int32(-32), v7)
	require.Equal(t, int64(-64), v8)
}

func TestC9(t *testing.T) {
	run := func() (string, uint8, int16, uint32, int64, float32, float64, complex64, complex128, error) {
		return "done", 255, -16, 32, -64, 3.14, 2.71, complex(5, 6), complex(7, 8), nil
	}

	v1, v2, v3, v4, v5, v6, v7, v8, v9 := must.C9(run())
	require.Equal(t, "done", v1)
	require.Equal(t, uint8(255), v2)
	require.Equal(t, int16(-16), v3)
	require.Equal(t, uint32(32), v4)
	require.Equal(t, int64(-64), v5)
	require.Equal(t, float32(3.14), v6)
	require.Equal(t, 2.71, v7)
	require.Equal(t, complex64(complex(5, 6)), v8)
	require.Equal(t, complex(7, 8), v9)
}
