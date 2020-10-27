package integer

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInteger(t *testing.T) {
	t.Parallel()

	t.Log("정수형 관련 유틸 함수 검증")

	t.Run("Min", func(t *testing.T) {
		const expected = 1

		t.Run("함수는 주어진 정수 중 가장 작은 수를 반환한다.", func(t *testing.T) {
			assert.Equal(t, MinInt(1, 3, 5), expected)
			assert.Equal(t, MinInt(3, 1, 5), expected)
			assert.Equal(t, MinInt(5, 3, 1), expected)
		})

		t.Run("32 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func(t *testing.T) {
			assert.Equal(t, MinInt32(1, 3, 5), int32(expected))
			assert.Equal(t, MinInt32(3, 1, 5), int32(expected))
			assert.Equal(t, MinInt32(5, 3, 1), int32(expected))
		})

		t.Run("64 함수는 주어진 정수 중 가장 작은 수를 반환한다.", func(t *testing.T) {
			assert.Equal(t, MinInt64(1, 3, 5), int64(expected))
			assert.Equal(t, MinInt64(3, 1, 5), int64(expected))
			assert.Equal(t, MinInt64(5, 3, 1), int64(expected))
		})
	})

	t.Run("Max", func(t *testing.T) {
		const expected = 5

		t.Run("함수는 주어진 정수 중 가장 큰 수를 반환한다.", func(t *testing.T) {
			assert.Equal(t, MaxInt(1, 3, 5), expected)
			assert.Equal(t, MaxInt(3, 1, 5), expected)
			assert.Equal(t, MaxInt(5, 3, 1), expected)
		})

		t.Run("32 함수는 주어진 정수 중 가장 큰 수를 반환한다.", func(t *testing.T) {
			assert.Equal(t, MaxInt32(1, 3, 5), int32(expected))
			assert.Equal(t, MaxInt32(3, 1, 5), int32(expected))
			assert.Equal(t, MaxInt32(5, 3, 1), int32(expected))
		})

		t.Run("64 함수는 주어진 정수 중 가장 큰 수를 반환한다.", func(t *testing.T) {
			assert.Equal(t, MaxInt64(1, 3, 5), int64(expected))
			assert.Equal(t, MaxInt64(3, 1, 5), int64(expected))
			assert.Equal(t, MaxInt64(5, 3, 1), int64(expected))
		})
	})

	t.Run("Pow", func(t *testing.T) {
		t.Run("32 함수는 거듭 제곱 연산을 잘 수행한다.", func(t *testing.T) {
			for _, data := range []struct {
				n, p, expected int32
			}{
				{3, 5, int32(math.Pow(3, 5))},
				{0, 5, 0},
				{10, 0, 1},
			} {
				actual := PowInt32(data.n, data.p)
				assert.Equal(t, actual, data.expected)
			}
		})

		t.Run("64 함수는 거듭 제곱 연산을 잘 수행한다.", func(t *testing.T) {
			for _, data := range []struct {
				n, p, expected int64
			}{
				{3, 5, int64(math.Pow(3, 5))},
				{0, 5, 0},
				{10, 0, 1},
			} {
				actual := PowInt64(data.n, data.p)
				assert.Equal(t, actual, data.expected)
			}
		})
	})
}
