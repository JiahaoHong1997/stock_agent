package stock_enums

type TimeLevel string

const (
	TimeLevel_1Minute  TimeLevel = "1"  // 1分钟
	TimeLevel_5Minute  TimeLevel = "5"  // 5分钟
	TimeLevel_15Minute TimeLevel = "15" // 15分钟
	TimeLevel_30Minute TimeLevel = "30" // 30分钟
	TimeLevel_1Hour    TimeLevel = "60" // 1小时
	TimeLevel_1Day     TimeLevel = "d"  // 日线
	TimeLevel_1Week    TimeLevel = "w"  // 周线
	TimeLevel_1Month   TimeLevel = "m"  // 月线
	TimeLevel_1Year    TimeLevel = "y"  // 年线
)

func (t TimeLevel) String() string {
	return string(t)
}
