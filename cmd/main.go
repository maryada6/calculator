package main

import (
	"calculator/pkg/calculator"
	"calculator/pkg/handler"
	"calculator/pkg/parser"
	"calculator/pkg/reader"
	"fmt"
	"os"
)

var newCalculator = calculator.NewCalculator()

func init() {
	handler.RegisterHandler("add", newCalculator.Add)
	handler.RegisterHandler("subtract", newCalculator.Subtract)
	handler.RegisterHandler("divide", newCalculator.Divide)
	handler.RegisterHandler("multiply", newCalculator.Multiply)
	handler.RegisterHandler("cancel", newCalculator.Cancel)
	handler.RegisterHandler("exit", newCalculator.Exit)
}

func main() {
	loop := true
	for loop == true {
		fmt.Println(">")
		input := reader.Reader(os.Stdin)
		operation, value := parser.ParseInput(input)
		handler.ExecuteHandler(handler.Operation(operation), value)
		fmt.Println(newCalculator.GetCurrentValue())
	}
}
