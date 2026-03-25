package date

import (
	"time"
)

func GetTime() string {
	nowTime := time.Now()
	formattedTime := nowTime.Format("2006-01-02 15:04:05")

	return formattedTime
}