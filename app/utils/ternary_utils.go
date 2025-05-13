package utils

import "time"

func TernaryString(condition bool, trueValue string, falseValue string) string {
	if condition {
		return trueValue
	}
	return falseValue
}

func TernaryTime(condition bool, trueValue time.Duration, falseValue time.Duration) time.Duration {
	if condition {
		return trueValue
	}
	return falseValue
}
