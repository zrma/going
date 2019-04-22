package utils

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PrintTest 함수 검증", func() {
	It("정상 동작", func() {
		err := PrintTest(func() {
			fmt.Println("test")
		}, []string{
			"test",
		})
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("반환값 없음", func() {
		err := PrintTest(func() {
		}, []string{})
		Expect(err).ShouldNot(HaveOccurred())
	})

	It("반환값이 있어야 하는데 안 나오는 경우", func() {
		err := PrintTest(func() {
		}, []string{
			"failed",
		})
		Expect(err).Should(HaveOccurred())
	})

	It("반환값이 없어야 하는데 나오는 경우", func() {
		err := PrintTest(func() {
			fmt.Println("test")
		}, []string{})
		Expect(err).Should(HaveOccurred())
	})
})
