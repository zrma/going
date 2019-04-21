package integer

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math"
)

var _ = Describe("정수형 관련 유틸 함수 검증", func() {
	Context("Min", func() {
		It("32 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func() {
			Expect(MinInt32([]int32{1, 3, 5})).Should(BeNumerically("==", 1))
			Expect(MinInt32([]int32{3, 1, 5})).Should(BeNumerically("==", 1))
			Expect(MinInt32([]int32{5, 3, 1})).Should(BeNumerically("==", 1))
		})

		It("64 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func() {
			Expect(MinInt64([]int64{1, 3, 5})).Should(BeNumerically("==", 1))
			Expect(MinInt64([]int64{3, 1, 5})).Should(BeNumerically("==", 1))
			Expect(MinInt64([]int64{5, 3, 1})).Should(BeNumerically("==", 1))
		})
	})

	Context("Max", func() {
		It("32 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func() {
			Expect(MaxInt32([]int32{1, 3, 5})).Should(BeNumerically("==", 5))
			Expect(MaxInt32([]int32{3, 1, 5})).Should(BeNumerically("==", 5))
			Expect(MaxInt32([]int32{5, 3, 1})).Should(BeNumerically("==", 5))
		})

		It("64 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func() {
			Expect(MaxInt64([]int64{1, 3, 5})).Should(BeNumerically("==", 5))
			Expect(MaxInt64([]int64{3, 1, 5})).Should(BeNumerically("==", 5))
			Expect(MaxInt64([]int64{5, 3, 1})).Should(BeNumerically("==", 5))
		})
	})

	Context("Pow 함수는", func() {
		It("거듭제곱 연산을 잘 수행한다.", func() {
			actual := PowInt64(3, 5)
			expected := math.Pow(3, 5)
			Expect(actual).Should(BeNumerically("==", expected))

			actual = PowInt64(0, 5)
			Expect(actual).Should(BeNumerically("==", 0))

			actual = PowInt64(10, 0)
			Expect(actual).Should(BeNumerically("==", 1))
		})
	})
})
