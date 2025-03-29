package mustnum_test

import (
	"testing"

	"github.com/yyle88/must/internal/tests"
	"github.com/yyle88/must/mustnum"
)

func TestLess(t *testing.T) {
	mustnum.Less(100, 200)
	mustnum.Less(0.1, 0.2)

	tests.ExpectPanic(t, func() {
		mustnum.Less(2, 1)
	})

	tests.ExpectPanic(t, func() {
		mustnum.Less(0.90, 0.11)
	})
}

func TestLt(t *testing.T) {
	mustnum.Lt(100, 200)
	mustnum.Lt(0.1, 0.2)

	tests.ExpectPanic(t, func() {
		mustnum.Lt(2, 1)
	})

	tests.ExpectPanic(t, func() {
		mustnum.Lt(0.90, 0.11)
	})
}

func TestLte(t *testing.T) {
	mustnum.Lte(100, 200)
	mustnum.Lte(5, 5)
	mustnum.Lte(0.1, 0.2)

	tests.ExpectPanic(t, func() {
		mustnum.Lte(2, 1)
	})

	tests.ExpectPanic(t, func() {
		mustnum.Lte(0.90, 0.11)
	})
}

func TestGt(t *testing.T) {
	mustnum.Gt(200, 100)
	mustnum.Gt(0.2, 0.1)

	tests.ExpectPanic(t, func() {
		mustnum.Gt(1, 2)
	})

	tests.ExpectPanic(t, func() {
		mustnum.Gt(0.11, 0.90)
	})
}

func TestGte(t *testing.T) {
	mustnum.Gte(200, 100)
	mustnum.Gte(5, 5)
	mustnum.Gte(0.2, 0.1)

	tests.ExpectPanic(t, func() {
		mustnum.Gte(1, 2)
	})

	tests.ExpectPanic(t, func() {
		mustnum.Gte(0.11, 0.90)
	})
}
