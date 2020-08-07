package utils_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/gomega"

	"github.com/zrma/going/utils"
)

func TestPrintTest(t *testing.T) {
	g := NewWithT(t)

	t.Run("정상 동작", func(t *testing.T) {
		err := utils.PrintTest(func() {
			fmt.Println("test")
		}, []string{
			"test",
		})
		g.Expect(err).ShouldNot(HaveOccurred())
	})

	t.Run("반환값 없음", func(t *testing.T) {
		err := utils.PrintTest(func() {
		}, []string{})
		g.Expect(err).ShouldNot(HaveOccurred())
	})

	t.Run("반환값이 있어야 하는데 안 나오는 경우", func(t *testing.T) {
		err := utils.PrintTest(func() {
		}, []string{
			"failed",
		})
		g.Expect(err).Should(HaveOccurred())
	})

	t.Run("반환값이 없어야 하는데 나오는 경우", func(t *testing.T) {
		err := utils.PrintTest(func() {
			fmt.Println("test")
		}, []string{})
		g.Expect(err).Should(HaveOccurred())
	})
}
