package main

import (
	"calculator/pkg/calculator"
	"calculator/pkg/handler"
	"calculator/pkg/parser"
	"calculator/pkg/reader"
	"calculator/pkg/view"
	"os"
)

func main() {
	var newCalculator = calculator.NewCalculator()
	for true {
		view.InitialSign()
		parseInput := parser.ParseInput(reader.Reader(os.Stdin))
		handler.ExecuteHandler(parseInput.Operation, parseInput.Value, newCalculator)
	}
}
