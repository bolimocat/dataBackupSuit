package tool

import (
	"strconv"
)

func Timetomillisecond(time string) int {
	return strconv.FormatInt(time, 10)
}