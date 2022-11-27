package utils

import (
	"fmt"
	"time"
)

func AddHour(h int) string {
	// h = 24

	nowTime := time.Now()

	afterHour := fmt.Sprintf("+%dh", h)

	durationHour, _ := time.ParseDuration(afterHour)

	retFmtTime := nowTime.Add(durationHour).Format("2006-01-02 15:04:05")

	return retFmtTime

}
