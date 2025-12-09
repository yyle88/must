package mustboolean_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/must/mustboolean"
)

func TestTrue(t *testing.T) {
	mustboolean.True(true)

	require.Panics(t, func() {
		mustboolean.True(false)
	})
}

func TestConflict(t *testing.T) {
	// Zero true - should pass
	mustboolean.Conflict()
	mustboolean.Conflict(false)
	mustboolean.Conflict(false, false)
	mustboolean.Conflict(false, false, false)

	// One true - should pass
	mustboolean.Conflict(true)
	mustboolean.Conflict(true, false)
	mustboolean.Conflict(false, true)
	mustboolean.Conflict(false, true, false)
	mustboolean.Conflict(false, false, true)

	// Two true - should panic
	require.Panics(t, func() {
		mustboolean.Conflict(true, true)
	})

	require.Panics(t, func() {
		mustboolean.Conflict(true, false, true)
	})

	require.Panics(t, func() {
		mustboolean.Conflict(true, true, false)
	})

	require.Panics(t, func() {
		mustboolean.Conflict(false, true, true)
	})

	// Three true - should panic
	require.Panics(t, func() {
		mustboolean.Conflict(true, true, true)
	})
}
