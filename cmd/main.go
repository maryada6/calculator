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
	handler.InitHandler(newCalculator)
	for true {
		view.ViewInitalSign()
		operation, value := parser.ParseInput(reader.Reader(os.Stdin))
		handler.ExecuteHandler(handler.Operation(operation), value, newCalculator)
	}
}
