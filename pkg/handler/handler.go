package handler

type Operation string

const add Operation = "add"

type HandleFunction func(float642 float64)
type HandleMap map[Operation]HandleFunction

var hf = HandleMap{}

func RegisterHandler(operator Operation, function HandleFunction) {
	hf[operator] = function
}
func ExecuteHandler(operation Operation, value float64) {
	hf[operation](value)
}

//func Handler(operation string, value float64, icalculator calculator.Icalculator) {
//	mapFunction := map[string]func(float642 float64){
//		"add":      icalculator.Add,
//		"subtract": icalculator.Subtract,
//		"multiply": icalculator.Multiply,
//		"divide":   icalculator.Divide,
//	}
//	mapFunction[operation](value)
//	fmt.Println(icalculator.GetCurrentValue())
//}
