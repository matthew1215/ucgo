package iutil

import (
	"strconv"
	"time"
)

func GetTimeStamp() int {
	timeStamp, _ := strconv.Atoi(strconv.FormatInt(time.Now().UTC().UnixNano(), 10))
	return timeStamp
}
