package pkg

type Calculator struct {
	currentValue float64
}

func NewCalculator(value float64) Calculator {
	return Calculator{value}
}
