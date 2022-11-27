package utils

import (
	"strconv"
	"unsafe"

	"go-micro.dev/v4/logger"
)

func StrToInt32(str string) int32 {
	vInt64, _ := strconv.ParseInt(str, 10, 64)
	vInt := *(*int32)(unsafe.Pointer(&vInt64))
	return vInt
}

func StrToFloat32(str string) float32 {
	num64, err := strconv.ParseFloat(str, 32)
	if err != nil {
		logger.Error(err)
		return 0.0
	}
	return float32(num64)
}
