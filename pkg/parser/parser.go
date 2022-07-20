package parser

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var commandList = []string{"add", "subtract", "multiply", "divide", "cancel", "exit"}

func Parser(inputString string) []string {
	split := strings.Split(inputString, " ")
	return split
}

func stringInList(operation string, list []string) bool {
	for _, item := range list {
		if item == operation {
			return true
		}
	}
	return false
}

func Validator(tokens []string) bool {
	if len(tokens) > 2 {
		return false
	}
	if len(tokens) == 1 {
		operation := tokens[0]
		return operation == "exit" || operation == "cancel"
	}
	operation, value := tokens[0], tokens[1]
	_, err := strconv.ParseFloat(value, 64)
	return stringInList(operation, commandList) && err == nil
}

func ParseInput(inputString string) (string, float64) {
	tokens := Parser(inputString)
	if Validator(tokens) {
		if len(tokens) == 1 {
			return tokens[0], 0.00
		}
		value, _ := strconv.ParseFloat(tokens[1], 64)
		return tokens[0], value
	} else {
		fmt.Println("Invalid input")
		os.Exit(0)
	}
	return "", 0.00
}
