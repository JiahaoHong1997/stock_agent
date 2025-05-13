package utils

import (
	"reflect"
	"time"
)

func ToString(input interface{}) string {
	val := reflect.ValueOf(input)

	// 检查是否为 string 类型
	if val.Kind() == reflect.String {
		return val.String()
	}

	// 检查是否为指向 string 类型的指针
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.String {
		return val.Elem().String()
	}

	return ""
}

func ToDuration(input interface{}) time.Duration {
	val := reflect.ValueOf(input)

	// 检查是否为 time.Duration 类型
	if val.Kind() == reflect.Int64 {
		if duration, ok := val.Interface().(time.Duration); ok {
			return duration
		}
	}

	// 检查是否为指向 time.Duration 类型的指针
	if val.Kind() == reflect.Ptr && val.Elem().Kind() == reflect.Int64 {
		if duration, ok := val.Elem().Interface().(time.Duration); ok {
			return duration
		}
	}

	return time.Duration(0)
}
