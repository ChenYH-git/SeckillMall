package utils

import "encoding/json"

func MapToStr(m map[string]interface{}) string {
	byteData, _ := json.Marshal(m)
	return string(byteData)
}

func StrToMap(s string) map[string]interface{} {
	var m map[string]interface{}
	err := json.Unmarshal([]byte(s), &m)
	if err != nil {
		panic(err)
	}
	return m
}
