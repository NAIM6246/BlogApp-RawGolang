package utils

import "strconv"

func ParseInt(str string) int64 {
	n, _ := strconv.Atoi(str)
	return int64(n)
}
