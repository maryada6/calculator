package handler

import (
	"calculator/pkg/calculator"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdditionHandler(t *testing.T) {
	t.Run("should add 10 and result as 10", func(t *testing.T) {
		calc := calculator.NewCalculator()
		AdditionHandler(10, calc)
		assert.Equal(t, 10.0, calc.GetCurrentValue())
	})
}

func TestSubtractionHandlerHandler(t *testing.T) {
	t.Run("should subtract 10 and result as -10", func(t *testing.T) {
		calc := calculator.NewCalculator()
		SubtractionHandler(10, calc)
		assert.Equal(t, -10.0, calc.GetCurrentValue())
	})
}

func TestMultiplicationHandler(t *testing.T) {
	t.Run("should multiply 10 with 2 and result as 20", func(t *testing.T) {
		calc := calculator.NewCalculator()
		AdditionHandler(2, calc)
		MultiplicationHandler(10, calc)
		assert.Equal(t, 20.0, calc.GetCurrentValue())
	})
}

func TestDivisionHandler(t *testing.T) {
	t.Run("should divide 10 by 2 and result as 5", func(t *testing.T) {
		calc := calculator.NewCalculator()
		AdditionHandler(10, calc)
		DivisionHandler(2, calc)
		assert.Equal(t, 5.0, calc.GetCurrentValue())
	})
}

func TestCancelHandler(t *testing.T) {
	t.Run("should reset calculator value to 0", func(t *testing.T) {
		calc := calculator.NewCalculator()
		AdditionHandler(10, calc)
		CancelHandler(0, calc)
		assert.Equal(t, 0.0, calc.GetCurrentValue())
	})
}

func TestRegisterHandler(t *testing.T) {
	t.Run("should add add operation to the map", func(t *testing.T) {
		handleMap := HandleMap{}
		assert.Equal(t, 0, len(handleMap))
		operation := Operation("add")
		handleMap.RegisterHandler(operation, AdditionHandler)
		assert.Equal(t, 1, len(handleMap))
	})
}
