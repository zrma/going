package str

import (
	"sort"

	"github.com/go-test/deep"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("문자열 관련 유틸 함수 검증", func() {
	//noinspection SpellCheckingInspection
	Context("Sort 함수는", func() {
		It("하나의 문자열을 잘 정렬한다.", func() {
			actual := Sort("dcba")
			Expect(actual).Should(Equal("abcd"))

			actual = Sort("ffbbaa")
			Expect(actual).Should(Equal("aabbff"))
		})
	})

	Context("SortAdapter 구조체는", func() {
		actual := []string{"def", "abc", "bcd"}

		It("정상적으로 잘 정렬한다.", func() {
			expected := []string{"abc", "bcd", "def"}

			sort.Sort(SortAdapter(actual))
			Expect(deep.Equal(actual, expected)).Should(BeNil())
		})
	})

	Context("Reverse 함수는", func() {
		It("문자열을 잘 뒤집는다.", func() {
			actual := Reverse("abc")
			Expect(actual).Should(Equal("cba"))
		})

		It("빈 문자열에 오류가 발생하지 않는다.", func(done Done) {
			go func() {
				defer GinkgoRecover()

				actual := Reverse("")
				Expect(actual).Should(Equal(""))

				close(done)
			}()
		})
	})
})
