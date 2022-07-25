package handler

import (
	"calculator/pkg/calculator"
	"calculator/pkg/view"
)

type Operation string

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

func NoHandler(value float64, icalculator calculator.Icalculator) {
	view.InvalidInput("Invalid input")
}

func init() {
	hf.RegisterHandler(add, AdditionHandler)
	hf.RegisterHandler(subtract, SubtractionHandler)
	hf.RegisterHandler(divide, DivisionHandler)
	hf.RegisterHandler(multiply, MultiplicationHandler)
	hf.RegisterHandler(cancel, CancelHandler)
	hf.RegisterHandler(exit, ExitHandler)
}

func (hf HandleMap) RegisterHandler(operator Operation, function HandleFunction) {
	hf[operator] = function
}

func ExecuteHandler(operation Operation, value float64, icalculator calculator.Icalculator) {
	_, isPresent := hf[operation]
	if isPresent {
		hf[operation](value, icalculator)
		return
	}
	NoHandler(0.0, icalculator)
}
