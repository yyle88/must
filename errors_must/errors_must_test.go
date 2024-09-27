package errors_must_test

import (
	"testing"

	"github.com/pkg/errors"
	"github.com/yyle88/must/errors_must"
	"github.com/yyle88/must/internal/utils"
)

func TestDone(t *testing.T) {
	errors_must.Done(error(nil))

	utils.ExpectPanic(t, func() {
		errors_must.Done(errors.New("wa"))
	})
}
