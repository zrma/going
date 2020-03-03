package integer

import (
	"math"
	"testing"

	"gotest.tools/assert"
)

func TestMinMax(t *testing.T) {
	t.Run("정수형 관련 유틸 함수 검증", func(t *testing.T) {
		t.Parallel()

		t.Run("Min", func(t *testing.T) {
			t.Run("32 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func(t *testing.T) {
				assert.Equal(t, MinInt32(1, 3, 5), int32(1))
				assert.Equal(t, MinInt32(3, 1, 5), int32(1))
				assert.Equal(t, MinInt32(5, 3, 1), int32(1))
			})

			t.Run("64 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func(t *testing.T) {
				assert.Equal(t, MinInt64(1, 3, 5), int64(1))
				assert.Equal(t, MinInt64(3, 1, 5), int64(1))
				assert.Equal(t, MinInt64(5, 3, 1), int64(1))
			})
		})

		t.Run("Max", func(t *testing.T) {
			t.Run("32 함수는 주어진 정수 중 가장 큰 수를 반환한다.", func(t *testing.T) {
				assert.Equal(t, MaxInt32(1, 3, 5), int32(5))
				assert.Equal(t, MaxInt32(3, 1, 5), int32(5))
				assert.Equal(t, MaxInt32(5, 3, 1), int32(5))
			})

			t.Run("64 함수는 주어진 정수 중 가장 큰 수를 반환한다.", func(t *testing.T) {
				assert.Equal(t, MaxInt64(1, 3, 5), int64(5))
				assert.Equal(t, MaxInt64(3, 1, 5), int64(5))
				assert.Equal(t, MaxInt64(5, 3, 1), int64(5))
			})
		})

		t.Run("Pow", func(t *testing.T) {
			t.Run("32 함수는 거듭제곱 연산을 잘 수행한다.", func(t *testing.T) {
				actual := PowInt32(3, 5)
				expected := math.Pow(3, 5)
				assert.Equal(t, actual, int32(expected))

				actual = PowInt32(0, 5)
				assert.Equal(t, actual, int32(0))

				actual = PowInt32(10, 0)
				assert.Equal(t, actual, int32(1))
			})

			t.Run("64 함수는 거듭제곱 연산을 잘 수행한다.", func(t *testing.T) {
				actual := PowInt64(3, 5)
				expected := math.Pow(3, 5)
				assert.Equal(t, actual, int64(expected))

				actual = PowInt64(0, 5)
				assert.Equal(t, actual, int64(0))

				actual = PowInt64(10, 0)
				assert.Equal(t, actual, int64(1))
			})
		})

	})
}
