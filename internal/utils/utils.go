package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func ExpectPanic(t *testing.T, run func()) {
	defer func() {
		if eco := recover(); eco != nil {
			t.Logf("expect panic then catch panic [%v] -> [success]", eco)
		} else {
			require.Fail(t, "expect panic while not panic -> [failure]")
		}
	}()

	run()
}
