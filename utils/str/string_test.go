package str

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStrUtils(t *testing.T) {
	//noinspection SpellCheckingInspection
	t.Run("문자열 관련 유틸 함수 검증", func(t *testing.T) {
		t.Parallel()

		t.Run("Sort 함수는 하나의 문자열을 잘 정렬한다.", func(t *testing.T) {
			actual := Sort("dcba")
			assert.Equal(t, actual, "abcd")

			actual = Sort("ffbbaa")
			assert.Equal(t, actual, "aabbff")
		})

		t.Run("SortAdapter 구조체가 정상적으로 잘 정렬한다.", func(t *testing.T) {
			actual := []string{"def", "abc", "bcd"}
			expected := []string{"abc", "bcd", "def"}

			sort.Sort(SortAdapter(actual))
			assert.Equal(t, actual, expected)
		})

		t.Run("Reverse 함수는", func(t *testing.T) {
			t.Run("문자열을 잘 뒤집는다.", func(t *testing.T) {
				actual := Reverse("abc")
				assert.Equal(t, actual, "cba")
			})

			t.Run("빈 문자열에 오류가 발생하지 않는다.", func(t *testing.T) {
				actual := Reverse("")
				assert.Equal(t, actual, "")
			})
		})
	})
}
