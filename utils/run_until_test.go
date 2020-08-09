package utils_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/zrma/going/utils"
)

var _ = Describe("RunUntil 테스트", func() {
	t := GinkgoT()
	It("성공", func(done Done) {
		defer close(done)
		RunUntil(t, func(holder Holder) {
			defer holder.Done()
		}, 1)
	}, 2)

	It("실패 - 타임아웃", func(done Done) {
		defer close(done)

		var newT tImpl
		RunUntil(&newT, func(holder Holder) {
			defer holder.Done()
			time.Sleep(3 * time.Second)
		}, 1)

		Expect(newT.failed()).Should(BeTrue())
	}, 2)
})

type tImpl struct {
	failedFlag bool
}

func (t *tImpl) Fatal(args ...interface{}) {
	t.failedFlag = true
}

func (t *tImpl) failed() bool {
	return t.failedFlag
}
