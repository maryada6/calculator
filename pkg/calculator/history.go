package calculator

type input struct {
	operation string
	value     float64
}

var history []input

func (calculator *Calculator) SaveInput(operation string, value float64) {
	history = append(history, input{operation, value})
}
