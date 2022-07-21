package parser

import (
	"calculator/pkg/handler"
	"strconv"
	"strings"
)

type Parser struct {
	Operation handler.Operation
	Value     float64
}

func NewParser(operation handler.Operation, float642 float64) Parser {
	return Parser{operation, float642}
}

func Split(inputString string) []string {
	split := strings.Split(inputString, " ")
	return split
}

func Validator(tokens []string) bool {
	return len(tokens) <= 2
}

func ParseInput(inputString string) Parser {
	tokens := Split(inputString)
	if Validator(tokens) {
		if len(tokens) == 1 {
			return NewParser(handler.Operation(tokens[0]), 0.00)
		}
		value, err := strconv.ParseFloat(tokens[1], 64)
		if err != nil {
			return NewParser(handler.Operation(tokens[0]), value)
		}
	}
	return NewParser(handler.Operation(""), 0.00)
}
