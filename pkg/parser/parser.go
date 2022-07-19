package parser

import (
	"strings"
)

func Parser(inputString string) []string {
	split := strings.Split(inputString, " ")
	return split
}
