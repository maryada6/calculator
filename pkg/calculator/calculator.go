package calculator

import "os"

type Icalculator interface {
	Add(value float64)
	Subtract(value float64)
	Multiply(value float64)
	Divide(value float64)
	GetCurrentValue() float64
	SetCurrentValue(value float64)
	Cancel(value float64)
	Exit(value float64)
}

type Calculator struct {
	currentValue float64
}

func (calculator *Calculator) SetCurrentValue(value float64) {
	calculator.currentValue = value
}

func (calculator *Calculator) GetCurrentValue() float64 {
	return calculator.currentValue
}

func NewCalculator() *Calculator {
	return &Calculator{0}
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
		panic("we cannot divide number by zero")
	}
	calculator.currentValue /= value
}

func (calculator *Calculator) Cancel(value float64) {
	calculator.currentValue = value
}

func (calculator *Calculator) Exit(value float64) {
	os.Exit(5)
}
