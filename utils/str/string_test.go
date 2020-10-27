package str

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStr(t *testing.T) {
	t.Parallel()

	t.Log("문자열 관련 유틸 함수 검증")

	//noinspection SpellCheckingInspection
	t.Run("Sort 함수는 하나의 문자열을 잘 정렬한다.", func(t *testing.T) {
		for _, data := range []struct {
			input, expected string
		}{
			{"dcba", "abcd"},
			{"ffbbaa", "aabbff"},
		} {
			actual := Sort(data.input)
			assert.Equal(t, actual, data.expected)
		}
	})

	t.Run("Reverse 함수는", func(t *testing.T) {
		t.Run("문자열을 잘 뒤집는다", func(t *testing.T) {
			actual := Reverse("abc")
			assert.Equal(t, actual, "cba")
		})

		t.Run("빈 문자열에 오류가 발생하지 않는다.", func(t *testing.T) {
			assert.NotPanics(t, func() {
				actual := Reverse("")
				assert.Equal(t, actual, "")
			})
		})
	})
}
