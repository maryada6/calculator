package handler

import (
	"calculator/pkg/calculator"
	"calculator/pkg/view"
)

type Operation string

const add Operation = "add"
const subtract Operation = "subtract"
const multiply Operation = "multiply"
const divide Operation = "divide"
const cancel Operation = "cancel"
const exit Operation = "exit"

type HandleFunction func(float642 float64, icalculator calculator.Icalculator)
type HandleMap map[Operation]HandleFunction

var hf = HandleMap{}

func AdditionHandler(value float64, icalculator calculator.Icalculator) {
	icalculator.Add(value)
	view.View(icalculator.GetCurrentValue())
}

func SubtractionHandler(value float64, icalculator calculator.Icalculator) {
	icalculator.Subtract(value)
	view.View(icalculator.GetCurrentValue())
}

func MultiplicationHandler(value float64, icalculator calculator.Icalculator) {
	icalculator.Multiply(value)
	view.View(icalculator.GetCurrentValue())
}

func DivisionHandler(value float64, icalculator calculator.Icalculator) {
	icalculator.Divide(value)
	if value != 0 {
		view.View(icalculator.GetCurrentValue())
	}
}

func CancelHandler(value float64, icalculator calculator.Icalculator) {
	icalculator.Cancel()
}

func ExitHandler(value float64, icalculator calculator.Icalculator) {
	icalculator.Exit()
}

func init() {
	RegisterHandler(add, AdditionHandler)
	RegisterHandler(subtract, SubtractionHandler)
	RegisterHandler(divide, DivisionHandler)
	RegisterHandler(multiply, MultiplicationHandler)
	RegisterHandler(cancel, CancelHandler)
	RegisterHandler(exit, ExitHandler)
}

func RegisterHandler(operator Operation, function HandleFunction) {
	hf[operator] = function
}

func ExecuteHandler(operation Operation, value float64, icalculator calculator.Icalculator) {
	hf[operation](value, icalculator)
}
