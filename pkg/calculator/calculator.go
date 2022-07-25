package calculator

import (
	"calculator/pkg/view"
	"os"
)

type Arithmetic interface {
	Add(value float64)
	Subtract(value float64)
	Multiply(value float64)
	Divide(value float64)
}

type Utility interface {
	Cancel()
	Exit()
}

type Icalculator interface {
	Arithmetic
	Utility
	GetCurrentValue() float64
	SaveInput(operation string, float642 float64)
}

type input struct {
	operation string
	value     float64
}

type Calculator struct {
	currentValue float64
	history      []input
}

func (calculator *Calculator) GetCurrentValue() float64 {
	return calculator.currentValue
}

func (calculator *Calculator) GetHistory() []input {
	return calculator.history
}

func NewCalculator() *Calculator {
	var inputs []input
	return &Calculator{0, inputs}
}

func (calculator *Calculator) Add(value float64) {
	calculator.currentValue += value
}

func (calculator *Calculator) Subtract(value float64) {
	calculator.currentValue -= value
}

func (calculator *Calculator) Multiply(value float64) {
	calculator.currentValue *= value
}

func (calculator *Calculator) Divide(value float64) {
	if value == 0.00 {
		view.InvalidInput("we cannot divide number by zero")
	} else {
		calculator.currentValue /= value

	}
}

func (calculator *Calculator) Cancel() {
	calculator.currentValue = 0.0
}

func (calculator *Calculator) Exit() {
	os.Exit(0)
}

func (calculator *Calculator) SaveInput(operation string, value float64) {
	calculator.history = append(calculator.history, input{operation, value})
}
