package stock_indicators

import (
	"math"
)

//------------------ 2级：技术指标函数(全部通过0级，1级函数实现） ------------------------------

// MACD 指标
func MACD(CLOSE []float64, SHORT, LONG, M int) ([]float64, []float64, []float64) {
	DIF := make([]float64, len(CLOSE))
	DEA := make([]float64, len(CLOSE))
	MACD := make([]float64, len(CLOSE))

	emaShort := EMA(CLOSE, SHORT)
	emaLong := EMA(CLOSE, LONG)

	for i := 0; i < len(CLOSE); i++ {
		DIF[i] = emaShort[i] - emaLong[i]
	}

	DEA = EMA(DIF, M)

	for i := 0; i < len(CLOSE); i++ {
		MACD[i] = (DIF[i] - DEA[i]) * 2
	}

	return DIF, DEA, MACD
}

// KDJ 指标
func KDJ(CLOSE, HIGH, LOW []float64, N, M1, M2 int) ([]float64, []float64, []float64) {
	RSV := make([]float64, len(CLOSE))
	K := make([]float64, len(CLOSE))
	D := make([]float64, len(CLOSE))
	J := make([]float64, len(CLOSE))

	llv := LLV(LOW, N)
	hhv := HHV(HIGH, N)

	for i := 0; i < len(CLOSE); i++ {
		if i < N-1 {
			RSV[i] = math.NaN()
			continue
		}
		if hhv[i]-llv[i] != 0 {
			RSV[i] = (CLOSE[i] - llv[i]) / (hhv[i] - llv[i]) * 100
		} else {
			RSV[i] = 0
		}
	}

	K = EMA(RSV, (M1*2)-1)
	D = EMA(K, (M2*2)-1)

	for i := 0; i < len(CLOSE); i++ {
		J[i] = K[i]*3 - D[i]*2
	}

	return K, D, J
}

// RSI 指标
func RSI(CLOSE []float64, N int) []float64 {
	DIF := make([]float64, len(CLOSE))
	for i := 1; i < len(CLOSE); i++ {
		DIF[i] = CLOSE[i] - CLOSE[i-1]
	}

	up := make([]float64, len(CLOSE))
	down := make([]float64, len(CLOSE))
	for i := 1; i < len(CLOSE); i++ {
		if DIF[i] > 0 {
			up[i] = DIF[i]
		} else {
			down[i] = -DIF[i]
		}
	}

	upMA := SMA(up, N, 1)
	downMA := SMA(down, N, 1)

	RSI := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		if downMA[i] != 0 {
			RSI[i] = upMA[i] / downMA[i] * 100
		} else {
			RSI[i] = 100
		}
	}

	return RSI
}

// WR 威廉指标
func WR(CLOSE, HIGH, LOW []float64, N, N1 int) ([]float64, []float64) {
	WR := make([]float64, len(CLOSE))
	WR1 := make([]float64, len(CLOSE))

	hhv := HHV(HIGH, N)
	llv := LLV(LOW, N)

	hhv1 := HHV(HIGH, N1)
	llv1 := LLV(LOW, N1)

	for i := 0; i < len(CLOSE); i++ {
		if i < N-1 {
			WR[i] = math.NaN()
		} else if hhv[i]-llv[i] != 0 {
			WR[i] = (hhv[i] - CLOSE[i]) / (hhv[i] - llv[i]) * 100
		} else {
			WR[i] = 0
		}

		if i < N1-1 {
			WR1[i] = math.NaN()
		} else if hhv1[i]-llv1[i] != 0 {
			WR1[i] = (hhv1[i] - CLOSE[i]) / (hhv1[i] - llv1[i]) * 100
		} else {
			WR1[i] = 0
		}
	}

	return WR, WR1
}

// BIAS 乖离率
func BIAS(CLOSE []float64, L1, L2, L3 int) ([]float64, []float64, []float64) {
	BIAS1 := make([]float64, len(CLOSE))
	BIAS2 := make([]float64, len(CLOSE))
	BIAS3 := make([]float64, len(CLOSE))

	ma1 := MA(CLOSE, L1)
	ma2 := MA(CLOSE, L2)
	ma3 := MA(CLOSE, L3)

	for i := 0; i < len(CLOSE); i++ {
		if i < L1-1 {
			BIAS1[i] = math.NaN()
		} else if ma1[i] != 0 {
			BIAS1[i] = (CLOSE[i] - ma1[i]) / ma1[i] * 100
		} else {
			BIAS1[i] = 0
		}

		if i < L2-1 {
			BIAS2[i] = math.NaN()
		} else if ma2[i] != 0 {
			BIAS2[i] = (CLOSE[i] - ma2[i]) / ma2[i] * 100
		} else {
			BIAS2[i] = 0
		}

		if i < L3-1 {
			BIAS3[i] = math.NaN()
		} else if ma3[i] != 0 {
			BIAS3[i] = (CLOSE[i] - ma3[i]) / ma3[i] * 100
		} else {
			BIAS3[i] = 0
		}
	}

	return BIAS1, BIAS2, BIAS3
}

// BOLL 布林带指标
func BOLL(CLOSE []float64, N int, P float64) ([]float64, []float64, []float64) {
	MID := MA(CLOSE, N)
	UPPER := make([]float64, len(CLOSE))
	LOWER := make([]float64, len(CLOSE))

	std := STD(CLOSE, N)

	for i := 0; i < len(CLOSE); i++ {
		if i < N-1 {
			UPPER[i] = math.NaN()
			LOWER[i] = math.NaN()
		} else {
			UPPER[i] = MID[i] + std[i]*P
			LOWER[i] = MID[i] - std[i]*P
		}
	}

	return UPPER, MID, LOWER
}

// PSY 心理线指标
func PSY(CLOSE []float64, N, M int) ([]float64, []float64) {
	PSY := make([]float64, len(CLOSE))
	PSYMA := make([]float64, len(CLOSE))

	ref := REF(CLOSE, 1)
	boolArr := make([]bool, len(CLOSE))
	for i := 1; i < len(CLOSE); i++ {
		boolArr[i] = CLOSE[i] > ref[i]
	}

	count := COUNT(boolArr, N)
	for i := 0; i < len(CLOSE); i++ {
		if i < N-1 {
			PSY[i] = math.NaN()
		} else {
			PSY[i] = count[i] / float64(N) * 100
		}
	}

	PSYMA = MA(PSY, M)

	return PSY, PSYMA
}

// CCI 商品通道指标
func CCI(CLOSE, HIGH, LOW []float64, N int) []float64 {
	TP := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		TP[i] = (HIGH[i] + LOW[i] + CLOSE[i]) / 3
	}

	maTP := MA(TP, N)
	avedevTP := AVEDEV(TP, N)

	CCI := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		if i < N-1 {
			CCI[i] = math.NaN()
		} else if avedevTP[i] != 0 {
			CCI[i] = (TP[i] - maTP[i]) / (0.015 * avedevTP[i])
		} else {
			CCI[i] = 0
		}
	}

	return CCI
}

// ATR 真实波动幅度平均值
func ATR(CLOSE, HIGH, LOW []float64, N int) []float64 {
	TR := make([]float64, len(CLOSE))

	// 第一个TR就是HIGH-LOW
	TR[0] = HIGH[0] - LOW[0]

	for i := 1; i < len(CLOSE); i++ {
		hl := HIGH[i] - LOW[i]
		hc := math.Abs(HIGH[i] - CLOSE[i-1])
		lc := math.Abs(LOW[i] - CLOSE[i-1])
		TR[i] = math.Max(math.Max(hl, hc), lc)
	}

	return MA(TR, N)
}

// BBI 多空指标
func BBI(CLOSE []float64, M1, M2, M3, M4 int) []float64 {
	BBI := make([]float64, len(CLOSE))

	ma1 := MA(CLOSE, M1)
	ma2 := MA(CLOSE, M2)
	ma3 := MA(CLOSE, M3)
	ma4 := MA(CLOSE, M4)

	for i := 0; i < len(CLOSE); i++ {
		// 只有当所有MA都有效时才计算
		if i >= M1-1 && i >= M2-1 && i >= M3-1 && i >= M4-1 {
			BBI[i] = (ma1[i] + ma2[i] + ma3[i] + ma4[i]) / 4
		} else {
			BBI[i] = math.NaN()
		}
	}

	return BBI
}

// DMI 动向指标
func DMI(CLOSE, HIGH, LOW []float64, M1, M2 int) ([]float64, []float64, []float64, []float64) {
	TR := make([]float64, len(CLOSE))
	HD := make([]float64, len(CLOSE))
	LD := make([]float64, len(CLOSE))

	// 计算TR, HD, LD
	for i := 1; i < len(CLOSE); i++ {
		TR[i] = math.Max(math.Max(HIGH[i]-LOW[i], math.Abs(HIGH[i]-CLOSE[i-1])), math.Abs(LOW[i]-CLOSE[i-1]))
		HD[i] = HIGH[i] - HIGH[i-1]
		LD[i] = LOW[i-1] - LOW[i]
	}

	// 计算DMP和DMM
	DMP := make([]float64, len(CLOSE))
	DMM := make([]float64, len(CLOSE))

	for i := 1; i < len(CLOSE); i++ {
		if HD[i] > 0 && HD[i] > LD[i] {
			DMP[i] = HD[i]
		} else {
			DMP[i] = 0
		}

		if LD[i] > 0 && LD[i] > HD[i] {
			DMM[i] = LD[i]
		} else {
			DMM[i] = 0
		}
	}

	// 计算M1日总和
	sumTR := SUM(TR, M1)
	sumDMP := SUM(DMP, M1)
	sumDMM := SUM(DMM, M1)

	// 计算PDI和MDI
	PDI := make([]float64, len(CLOSE))
	MDI := make([]float64, len(CLOSE))

	for i := 0; i < len(CLOSE); i++ {
		if i < M1-1 {
			PDI[i] = math.NaN()
			MDI[i] = math.NaN()
		} else if sumTR[i] != 0 {
			PDI[i] = sumDMP[i] / sumTR[i] * 100
			MDI[i] = sumDMM[i] / sumTR[i] * 100
		} else {
			PDI[i] = 0
			MDI[i] = 0
		}
	}

	// 计算ADX
	DX := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		if i < M1-1 {
			DX[i] = math.NaN()
		} else if (PDI[i] + MDI[i]) != 0 {
			DX[i] = math.Abs(PDI[i]-MDI[i]) / (PDI[i] + MDI[i]) * 100
		} else {
			DX[i] = 0
		}
	}

	ADX := MA(DX, M2)

	// 计算ADXR
	ADXR := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		if i < M1+M2-2 {
			ADXR[i] = math.NaN()
		} else {
			ADXR[i] = (ADX[i] + ADX[i-M2]) / 2
		}
	}

	return PDI, MDI, ADX, ADXR
}

// TAQ 唐安奇通道
func TAQ(HIGH, LOW []float64, N int) ([]float64, []float64, []float64) {
	UP := HHV(HIGH, N)
	DOWN := LLV(LOW, N)
	MID := make([]float64, len(HIGH))

	for i := 0; i < len(HIGH); i++ {
		if i < N-1 {
			MID[i] = math.NaN()
		} else {
			MID[i] = (UP[i] + DOWN[i]) / 2
		}
	}

	return UP, MID, DOWN
}

// TRIX 三重指数平滑平均线
func TRIX(CLOSE []float64, M1, M2 int) ([]float64, []float64) {
	// 三重EMA
	ema1 := EMA(CLOSE, M1)
	ema2 := EMA(ema1, M1)
	ema3 := EMA(ema2, M1)

	TRIX := make([]float64, len(CLOSE))
	for i := 1; i < len(CLOSE); i++ {
		if ema3[i-1] != 0 {
			TRIX[i] = (ema3[i] - ema3[i-1]) / ema3[i-1] * 100
		} else {
			TRIX[i] = 0
		}
	}

	TRMA := MA(TRIX, M2)

	return TRIX, TRMA
}

// VR 容量比率
func VR(CLOSE, VOL []float64, M1 int) []float64 {
	LC := REF(CLOSE, 1)

	upSum := make([]float64, len(CLOSE))
	downSum := make([]float64, len(CLOSE))

	for i := 1; i < len(CLOSE); i++ {
		if CLOSE[i] > LC[i] {
			upSum[i] = VOL[i]
		} else {
			downSum[i] = VOL[i]
		}
	}

	sumUp := SUM(upSum, M1)
	sumDown := SUM(downSum, M1)

	result := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		if i < M1-1 {
			result[i] = math.NaN()
		} else if sumDown[i] != 0 {
			result[i] = sumUp[i] / sumDown[i] * 100
		} else {
			result[i] = 0
		}
	}

	return result
}

// EMV 简易波动指标
func EMV(HIGH, LOW, VOL []float64, N, M int) ([]float64, []float64) {
	volume := make([]float64, len(VOL))
	maVol := MA(VOL, N)
	for i := 0; i < len(VOL); i++ {
		if maVol[i] != 0 {
			volume[i] = maVol[i] / VOL[i]
		} else {
			volume[i] = 0
		}
	}

	mid := make([]float64, len(HIGH))
	refHL := REF(HIGH, 1)
	refLL := REF(LOW, 1)
	for i := 1; i < len(HIGH); i++ {
		if (HIGH[i] + LOW[i]) != 0 {
			mid[i] = 100 * (HIGH[i] + LOW[i] - (refHL[i] + refLL[i])) / (HIGH[i] + LOW[i])
		} else {
			mid[i] = 0
		}
	}

	hlDiff := make([]float64, len(HIGH))
	for i := 0; i < len(HIGH); i++ {
		hlDiff[i] = HIGH[i] - LOW[i]
	}
	maHLDiff := MA(hlDiff, N)

	emv := make([]float64, len(HIGH))
	for i := 0; i < len(HIGH); i++ {
		if i < N-1 {
			emv[i] = math.NaN()
		} else if maHLDiff[i] != 0 {
			emv[i] = (mid[i] * volume[i] * (HIGH[i] - LOW[i])) / maHLDiff[i]
		} else {
			emv[i] = 0
		}
	}

	maEMV := MA(emv, M)

	return emv, maEMV
}

// DPO 区间震荡线
func DPO(CLOSE []float64, M1, M2, M3 int) ([]float64, []float64) {
	ma := MA(CLOSE, M1)
	refMA := REF(ma, M2)

	DPO := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		if i < M1+M2-1 {
			DPO[i] = math.NaN()
		} else {
			DPO[i] = CLOSE[i] - refMA[i]
		}
	}

	MADPO := MA(DPO, M3)

	return DPO, MADPO
}

// BRAR 情绪指标
func BRAR(OPEN, CLOSE, HIGH, LOW []float64, M1 int) ([]float64, []float64) {
	AR := make([]float64, len(CLOSE))
	BR := make([]float64, len(CLOSE))

	sumHighOpen := make([]float64, len(CLOSE))
	sumOpenLow := make([]float64, len(CLOSE))
	sumHighRefClose := make([]float64, len(CLOSE))
	sumRefCloseLow := make([]float64, len(CLOSE))

	for i := 1; i < len(CLOSE); i++ {
		sumHighOpen[i] = HIGH[i] - OPEN[i]
		sumOpenLow[i] = OPEN[i] - LOW[i]

		if HIGH[i] > CLOSE[i-1] {
			sumHighRefClose[i] = HIGH[i] - CLOSE[i-1]
		}
		if CLOSE[i-1] > LOW[i] {
			sumRefCloseLow[i] = CLOSE[i-1] - LOW[i]
		}
	}

	sumHO := SUM(sumHighOpen, M1)
	sumOL := SUM(sumOpenLow, M1)
	sumHRC := SUM(sumHighRefClose, M1)
	sumRCL := SUM(sumRefCloseLow, M1)

	for i := 0; i < len(CLOSE); i++ {
		if i < M1-1 {
			AR[i] = math.NaN()
			BR[i] = math.NaN()
		} else if sumOL[i] != 0 {
			AR[i] = sumHO[i] / sumOL[i] * 100
		} else {
			AR[i] = 0
		}

		if sumRCL[i] != 0 {
			BR[i] = sumHRC[i] / sumRCL[i] * 100
		} else {
			BR[i] = 0
		}
	}

	return AR, BR
}

// DMA 平行线差指标
func DMA(CLOSE []float64, N1, N2, M int) ([]float64, []float64) {
	ma1 := MA(CLOSE, N1)
	ma2 := MA(CLOSE, N2)

	DIF := make([]float64, len(CLOSE))
	for i := 0; i < len(CLOSE); i++ {
		if i >= N1-1 && i >= N2-1 {
			DIF[i] = ma1[i] - ma2[i]
		} else {
			DIF[i] = math.NaN()
		}
	}

	DIFMA := MA(DIF, M)

	return DIF, DIFMA
}

// MTM 动量指标
func MTM(CLOSE []float64, N, M int) ([]float64, []float64) {
	mtm := make([]float64, len(CLOSE))
	for i := N; i < len(CLOSE); i++ {
		mtm[i] = CLOSE[i] - CLOSE[i-N]
	}

	mtmma := MA(mtm, M)

	return mtm, mtmma
}
