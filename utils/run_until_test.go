package utils_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	. "github.com/zrma/going/utils"
)

func TestRunUnit(t *testing.T) {
	t.Parallel()

	t.Log("RunUntil")

	t.Run("성공", func(t *testing.T) {
		RunUntil(t, func(holder Holder) {
			defer holder.Done()
		}, 1)
	})

	t.Run("실패 타임아웃", func(t *testing.T) {

		var newT tImpl
		RunUntil(&newT, func(holder Holder) {
			defer holder.Done()
			time.Sleep(3 * time.Second)
		}, 1)

		assert.True(t, newT.failed())
	})
}

type tImpl struct {
	failedFlag bool
}

func (t *tImpl) Fatal(args ...interface{}) {
	t.failedFlag = true
}

func (t *tImpl) failed() bool {
	return t.failedFlag
}
