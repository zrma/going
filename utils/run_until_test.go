package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRunUntil(t *testing.T) {
	t.Run("테스트 성공", func(t *testing.T) {
		RunUntil(t, func(done Done) {
			defer close(done)
		}, 1)
	})

	t.Run("테스트 실패 - 타임아웃", func(t *testing.T) {
		var newT tImpl
		RunUntil(&newT, func(done Done) {
			defer close(done)
			time.Sleep(5 * time.Second)
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
