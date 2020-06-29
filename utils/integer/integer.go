package integer

import "math"

// MinInt 함수는 인자로 주어진 int 가변 인수 중 가장 작은 값을 반환한다.
func MinInt(arr ...int) int {
	var min = math.MaxInt32
	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

// MinInt32 함수는 인자로 주어진 int32 가변 인수 중 가장 작은 값을 반환한다.
func MinInt32(arr ...int32) int32 {
	var min int32 = math.MaxInt32
	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

// MinInt64 함수는 인자로 주어진 int64 가변 인수 중 가장 작은 값을 반환한다.
func MinInt64(arr ...int64) int64 {
	var min int64 = math.MaxInt64
	for _, num := range arr {
		if min > num {
			min = num
		}
	}
	return min
}

// MaxInt 함수는 인자로 주어진 int 가변 인수 중 가장 큰 값을 반환한다.
func MaxInt(arr ...int) int {
	var max = math.MinInt32
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

// MaxInt32 함수는 인자로 주어진 int32 가변 인수 중 가장 큰 값을 반환한다.
func MaxInt32(arr ...int32) int32 {
	var max int32 = math.MinInt32
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

// MaxInt64 함수는 인자로 주어진 int64 가변 인수 중 가장 큰 값을 반환한다.
func MaxInt64(arr ...int64) int64 {
	var max int64 = math.MinInt64
	for _, num := range arr {
		if max < num {
			max = num
		}
	}
	return max
}

// PowInt32 함수는 인자로 주어진 n의 p 거듭제곱 값을 반환한다.
func PowInt32(n, p int32) int32 {
	if n == 0 {
		return 0
	}

	var result int32 = 1
	for ; p > 0; p-- {
		result *= n
	}

	return result
}

// PowInt64 함수는 인자로 주어진 n의 p 거듭제곱 값을 반환한다.
func PowInt64(n, p int64) int64 {
	if n == 0 {
		return 0
	}

	var result int64 = 1
	for ; p > 0; p-- {
		result *= n
	}

	return result
}
