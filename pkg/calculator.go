package pkg

type Calculator struct {
	currentValue float64
}

func NewCalculator() Calculator {
	return Calculator{0}
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
