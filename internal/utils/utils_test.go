package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestZero(t *testing.T) {
	require.Equal(t, 0, Zero[int]())
	require.Equal(t, "", Zero[string]())
	require.Equal(t, false, Zero[bool]())
}
