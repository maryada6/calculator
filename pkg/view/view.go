package view

import "fmt"

func View(value float64) {
	fmt.Println(value)
}

func InitialSign() {
	fmt.Print(">")
}

func InvalidInput(string2 string) {
	fmt.Println(string2)
}
