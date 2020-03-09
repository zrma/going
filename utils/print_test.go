package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintTest(t *testing.T) {
	t.Run("PrintTest 함수 검증", func(t *testing.T) {
		t.Parallel()

		t.Run("정상 동작", func(t *testing.T) {
			err := PrintTest(func() {
				fmt.Println("test")
			}, []string{
				"test",
			})
			assert.NoError(t, err)
		})

		t.Run("반환값 없음", func(t *testing.T) {
			err := PrintTest(func() {
			}, []string{})
			assert.NoError(t, err)
		})

		t.Run("반환값이 있어야 하는데 안 나오는 경우", func(t *testing.T) {
			err := PrintTest(func() {
			}, []string{
				"failed",
			})
			assert.Error(t, err, "slice[0]: <no value> != failed")
		})

		t.Run("반환값이 없어야 하는데 나오는 경우", func(t *testing.T) {
			err := PrintTest(func() {
				fmt.Println("test")
			}, []string{})
			assert.Error(t, err, "slice[0]: test != <no value>")
		})
	})
}
