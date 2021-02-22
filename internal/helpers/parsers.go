package helpers

import (
	"strconv"
)

func ParseUint(value string) (uint, error) {
	u64, err := strconv.ParseUint(value, 10, 32)
	if err != nil {
		return 0, err
	}
	wd := uint(u64)
	return wd, nil
}

func ParseFloat(value float64) string {
	str := strconv.FormatFloat(value, 'f', 0, 64)
	strFinal := string(str)
	return strFinal
}
