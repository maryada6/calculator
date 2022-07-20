package validator

import "strconv"

var commandList = []string{"add", "subtract", "multiply", "divide", "cancel", "exit"}

func stringInList(operation string, list []string) bool {
	for _, item := range list {
		if item == operation {
			return true
		}
	}
	return false
}

func Validator(tokens []string) bool {
	if len(tokens) != 2 {
		return false
	}
	operation, value := tokens[0], tokens[1]
	_, err := strconv.ParseFloat(value, 64)
	return stringInList(operation, commandList) && err == nil
}
