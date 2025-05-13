package utils

import (
	"encoding/json"
)

func ToJsonString(input interface{}) string {
	result, err := json.Marshal(input)
	if err != nil {
		return ""
	}
	return string(result)
}
