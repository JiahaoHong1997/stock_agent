package stock_indicators

import (
	"math"
)

//------------------ 0级：核心工具函数 --------------------------------------------

// RD 四舍五入取D位小数
func RD(value float64, D int) float64 {
	shift := math.Pow(10, float64(D))
	return math.Round(value*shift) / shift
}

// RET 返回序列倒数第N个值,默认返回最后一个
func RET(S []float64, N int) float64 {
	if N <= 0 || N > len(S) {
		N = 1
	}
	return S[len(S)-N]
}

// ABS 返回绝对值
func ABS(value float64) float64 {
	return math.Abs(value)
}

// MAX 返回两个数中的较大值
func MAX(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

// MIN 返回两个数中的较小值
func MIN(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

// MA 求序列的N日简单移动平均值
func MA(S []float64, N int) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N-1 {
			result[i] = math.NaN()
			continue
		}
		sum := 0.0
		for j := i - N + 1; j <= i; j++ {
			sum += S[j]
		}
		result[i] = sum / float64(N)
	}
	return result
}

// REF 对序列整体下移动N,返回序列(shift后会产生NAN)
func REF(S []float64, N int) []float64 {
	if N < 0 || N >= len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N {
			result[i] = math.NaN()
		} else {
			result[i] = S[i-N]
		}
	}
	return result
}

// DIFF 前一个值减后一个值,前面会产生nan
func DIFF(S []float64, N int) []float64 {
	if N <= 0 || N >= len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N {
			result[i] = math.NaN()
		} else {
			result[i] = S[i] - S[i-N]
		}
	}
	return result
}

// STD 求序列的N日标准差，返回序列
func STD(S []float64, N int) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N-1 {
			result[i] = math.NaN()
			continue
		}

		// 计算平均值
		sum := 0.0
		for j := i - N + 1; j <= i; j++ {
			sum += S[j]
		}
		mean := sum / float64(N)

		// 计算方差
		variance := 0.0
		for j := i - N + 1; j <= i; j++ {
			variance += math.Pow(S[j]-mean, 2)
		}

		// 标准差
		result[i] = math.Sqrt(variance / float64(N))
	}
	return result
}

// IF 序列布尔判断
func IF(condition bool, trueVal, falseVal float64) float64 {
	if condition {
		return trueVal
	}
	return falseVal
}

// SUM 对序列求N天累计和，返回序列
func SUM(S []float64, N int) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N-1 {
			result[i] = math.NaN()
			continue
		}
		sum := 0.0
		for j := i - N + 1; j <= i; j++ {
			sum += S[j]
		}
		result[i] = sum
	}
	return result
}

// HHV 最近N天的最大值
func HHV(S []float64, N int) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N-1 {
			result[i] = math.NaN()
			continue
		}
		max := S[i]
		for j := i - N + 1; j <= i; j++ {
			if S[j] > max {
				max = S[j]
			}
		}
		result[i] = max
	}
	return result
}

// LLV 最近N天的最小值
func LLV(S []float64, N int) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N-1 {
			result[i] = math.NaN()
			continue
		}
		min := S[i]
		for j := i - N + 1; j <= i; j++ {
			if S[j] < min {
				min = S[j]
			}
		}
		result[i] = min
	}
	return result
}

// EMA 指数移动平均
func EMA(S []float64, N int) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))
	k := 2.0 / float64(N+1)

	var startIndex int
	for i := 0; i < len(S); i++ {
		if math.IsNaN(S[i]) {
			result[i] = math.NaN()
			startIndex = i + 1
		}
	}
	// 第一个EMA值为第一个数据点
	result[startIndex] = S[startIndex]

	for i := startIndex + 1; i < len(S); i++ {
		result[i] = S[i]*k + result[i-1]*(1-k)
	}
	return result
}

// SMA 中国式的SMA
func SMA(S []float64, N int, M float64) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))

	// 前N个值为简单移动平均
	for i := 0; i < N && i < len(S); i++ {
		sum := 0.0
		for j := 0; j <= i; j++ {
			sum += S[j]
		}
		result[i] = sum / float64(i+1)
	}

	// 从N+1开始使用SMA公式
	for i := N; i < len(S); i++ {
		result[i] = (M*S[i] + (float64(N)-M)*result[i-1]) / float64(N)
	}
	return result
}

// AVEDEV 平均绝对偏差
func AVEDEV(S []float64, N int) []float64 {
	if N <= 0 || N > len(S) {
		return nil
	}

	result := make([]float64, len(S))
	for i := 0; i < len(S); i++ {
		if i < N-1 {
			result[i] = math.NaN()
			continue
		}

		// 计算平均值
		sum := 0.0
		for j := i - N + 1; j <= i; j++ {
			sum += S[j]
		}
		mean := sum / float64(N)

		// 计算平均绝对偏差
		devSum := 0.0
		for j := i - N + 1; j <= i; j++ {
			devSum += math.Abs(S[j] - mean)
		}
		result[i] = devSum / float64(N)
	}
	return result
}

// SLOPE 线性回归斜率
func SLOPE(S []float64, N int, returnSeries bool) (float64, []float64) {
	if N <= 0 || N > len(S) {
		return math.NaN(), nil
	}

	// 只取最后N个数据点
	M := S[len(S)-N:]

	// 计算线性回归
	var sumX, sumY, sumXY, sumX2 float64
	for i := 0; i < len(M); i++ {
		x := float64(i)
		y := M[i]
		sumX += x
		sumY += y
		sumXY += x * y
		sumX2 += x * x
	}

	n := float64(len(M))
	slope := (n*sumXY - sumX*sumY) / (n*sumX2 - sumX*sumX)
	intercept := (sumY - slope*sumX) / n

	if !returnSeries {
		return slope, nil
	}

	// 计算整个回归线
	Y := make([]float64, len(M))
	for i := range M {
		Y[i] = intercept + slope*float64(i)
	}

	return Y[len(Y)-1] - Y[0], Y
}
