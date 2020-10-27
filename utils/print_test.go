package utils_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/zrma/going/utils"
)

func TestPrintTest(t *testing.T) {
	t.Run("정상 동작", func(t *testing.T) {
		err := utils.PrintTest(func() {
			fmt.Println("test")
		}, []string{
			"test",
		})
		assert.NoError(t, err)
	})

	t.Run("반환값 없음", func(t *testing.T) {
		err := utils.PrintTest(func() {
		}, []string{})
		assert.NoError(t, err)
	})

	t.Run("반환값이 있어야 하는데 안 나오는 경우", func(t *testing.T) {
		err := utils.PrintTest(func() {
		}, []string{
			"failed",
		})
		assert.Error(t, err)
	})

	t.Run("반환값이 없어야 하는데 나오는 경우", func(t *testing.T) {
		err := utils.PrintTest(func() {
			fmt.Println("test")
		}, []string{})
		assert.Error(t, err)
	})
}
