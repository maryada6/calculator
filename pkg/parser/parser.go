package parser

import (
	"fmt"
)

func Reader() string {
	var inputString string
	fmt.Scanf("%s", &inputString)
	return inputString
}
