package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should initialise the calculator entity with value 0", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewCalculator()
		})
	})

	t.Run("Should return calculator entity", func(t *testing.T) {
		assert.IsType(t, Calculator{}, NewCalculator())
	})
}

func TestAdd(t *testing.T) {
	t.Run("should return same value on adding with 0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		assert.Equal(t, 5.00, calculator.currentValue)
	})

	t.Run("should return 10 on adding 5 to 5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		calculator.Add(5)
		assert.Equal(t, 10.00, calculator.currentValue)
	})

	t.Run("should return -5 on adding 5 to -10", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-10)
		calculator.Add(5)
		assert.Equal(t, -5.00, calculator.currentValue)
	})

	t.Run("should return -5 on adding -2 to -3", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-3)
		calculator.Add(-2)
		assert.Equal(t, -5.00, calculator.currentValue)
	})

	t.Run("should return 5.50 on adding 3.50 to 2", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(2)
		calculator.Add(3.50)
		assert.Equal(t, 5.50, calculator.currentValue)
	})
}

func TestSubtract(t *testing.T) {
	t.Run("should return -1 times the same value on subtracting with 0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Subtract(5)
		assert.Equal(t, -5.00, calculator.currentValue)
	})
}
