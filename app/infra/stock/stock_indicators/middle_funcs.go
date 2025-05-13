package stock_indicators

import "math"

//------------------ 1级：应用层函数(通过0级核心函数实现） ----------------------------------

// COUNT 最近N天满足条件的天数
func COUNT(S_BOOL []bool, N int) []float64 {
	result := make([]float64, len(S_BOOL))
	for i := 0; i < len(S_BOOL); i++ {
		if i < N-1 {
			result[i] = math.NaN()
			continue
		}
		count := 0
		for j := i - N + 1; j <= i; j++ {
			if S_BOOL[j] {
				count++
			}
		}
		result[i] = float64(count)
	}
	return result
}

// EVERY 最近N天是否都是True
func EVERY(S_BOOL []bool, N int) []bool {
	result := make([]bool, len(S_BOOL))
	for i := 0; i < len(S_BOOL); i++ {
		if i < N-1 {
			result[i] = false
			continue
		}
		allTrue := true
		for j := i - N + 1; j <= i; j++ {
			if !S_BOOL[j] {
				allTrue = false
				break
			}
		}
		result[i] = allTrue
	}
	return result
}

// LAST 从前A日到前B日一直满足条件
func LAST(S_BOOL []bool, A, B int) bool {
	if A < B {
		A = B
	}
	if A >= len(S_BOOL) {
		return false
	}

	start := len(S_BOOL) - A
	end := len(S_BOOL) - B
	if start < 0 || end > len(S_BOOL) {
		return false
	}

	for i := start; i < end; i++ {
		if !S_BOOL[i] {
			return false
		}
	}
	return true
}

// EXIST N日内是否存在一天满足条件
func EXIST(S_BOOL []bool, N int) bool {
	if N <= 0 || N > len(S_BOOL) {
		return false
	}

	start := len(S_BOOL) - N
	if start < 0 {
		start = 0
	}

	for i := start; i < len(S_BOOL); i++ {
		if S_BOOL[i] {
			return true
		}
	}
	return false
}

// BARSLAST 上一次条件成立到当前的周期
func BARSLAST(S_BOOL []bool) int {
	for i := len(S_BOOL) - 1; i >= 0; i-- {
		if S_BOOL[i] {
			return len(S_BOOL) - i - 1
		}
	}
	return -1
}

// FORCAST 线性回归后的预测值
func FORCAST(S []float64, N int) float64 {
	slope, Y := SLOPE(S, N, true)
	if len(Y) == 0 {
		return math.NaN()
	}
	return Y[len(Y)-1] + slope
}

// CROSS 判断两条线是否交叉
func CROSS(S1, S2 []float64) bool {
	if len(S1) < 2 || len(S2) < 2 {
		return false
	}

	prev1 := S1[len(S1)-2] <= S2[len(S2)-2]
	curr1 := S1[len(S1)-1] > S2[len(S2)-1]

	prev2 := S1[len(S1)-2] >= S2[len(S2)-2]
	curr2 := S1[len(S1)-1] < S2[len(S2)-1]

	return (prev1 && curr1) || (prev2 && curr2)
}
