package utils

import (
	"testing"

	"github.com/pkg/errors"
)

func TestExpectPanic(t *testing.T) {
	ExpectPanic(t, func() {
		panic(errors.New("exp"))
	})
}
