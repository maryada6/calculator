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
