package util

import "strconv"

func ToInteger64(str string) int64 {
	value, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		panic(err)
	}

	return value
}

func ToInteger(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return value
}
