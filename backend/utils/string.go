package utils

import (
	"strconv"
)

func IsInt(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}
