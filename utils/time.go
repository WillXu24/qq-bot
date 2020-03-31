package utils

import "time"

func GetTimeZero() int64 {
	t := time.Now()
	newTime := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return newTime.Unix()
}
