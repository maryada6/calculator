package parser

import (
	"strconv"
	"strings"
)

func Parser(inputString string) (string, float64) {
	split := strings.Split(inputString, " ")
	value, _ := strconv.ParseFloat(split[1], 64)
	return split[0], value
}
