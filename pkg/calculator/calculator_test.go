package calculator

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

	t.Run("should return 10.00 on adding 5 to 5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		calculator.Add(5)
		assert.Equal(t, 10.00, calculator.currentValue)
	})

	t.Run("should return -5.00 on adding 5 to -10", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-10)
		calculator.Add(5)
		assert.Equal(t, -5.00, calculator.currentValue)
	})

	t.Run("should return 5.00 on adding -5 to 10", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(10)
		calculator.Add(-5)
		assert.Equal(t, 5.00, calculator.currentValue)
	})

	t.Run("should return -5.00 on adding -2 to -3", func(t *testing.T) {
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

	t.Run("should return 0.00 on subtracting 5 from 5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		calculator.Subtract(5)
		assert.Equal(t, 0.00, calculator.currentValue)
	})

	t.Run("should return -11.00 on subtracting 5 from -6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-6)
		calculator.Subtract(5)
		assert.Equal(t, -11.00, calculator.currentValue)
	})

	t.Run("should return 11.00 on subtracting -5 from 6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(6)
		calculator.Subtract(-5)
		assert.Equal(t, 11.00, calculator.currentValue)
	})

	t.Run("should return -1.00 on subtracting -5 from -6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-6)
		calculator.Subtract(-5)
		assert.Equal(t, -1.00, calculator.currentValue)
	})

	t.Run("should return -0.50 on subtracting -5.50 from -6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-6)
		calculator.Subtract(-5.50)
		assert.Equal(t, -0.50, calculator.currentValue)
	})
}

func TestMultiply(t *testing.T) {
	t.Run("should return 0 on multiplication with 0", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(5)
		calculator.Multiply(0)
		assert.Equal(t, 0.00, calculator.currentValue)
	})

	t.Run("should return same value on multiplication with 1", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(8)
		calculator.Multiply(1)
		assert.Equal(t, 8.00, calculator.currentValue)
	})

	t.Run("should return 30.00  on multiplication of 5 with 6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(6)
		calculator.Multiply(5)
		assert.Equal(t, 30.00, calculator.currentValue)
	})

	t.Run("should return -30.00  on multiplication of -5 with 6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(6)
		calculator.Multiply(-5)
		assert.Equal(t, -30.00, calculator.currentValue)
	})

	t.Run("should return -30.00  on multiplication of 5 with -6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-6)
		calculator.Multiply(5)
		assert.Equal(t, -30.00, calculator.currentValue)
	})

	t.Run("should return 30.00  on multiplication of -5 with -6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-6)
		calculator.Multiply(-5)
		assert.Equal(t, 30.00, calculator.currentValue)
	})

	t.Run("should return 31.50  on multiplication of 5.25 with 6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(6)
		calculator.Multiply(5.25)
		assert.Equal(t, 31.50, calculator.currentValue)
	})
}

func TestDivide(t *testing.T) {
	t.Run("should not divide a value by 0", func(t *testing.T) {
		assert.Panics(t, func() {
			calculator := NewCalculator()
			calculator.Add(6)
			calculator.Divide(0)
		})
	})

	t.Run("should return same value on division by 1", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(6)
		calculator.Divide(1)
		assert.Equal(t, 6.00, calculator.currentValue)
	})

	t.Run("should return 5.00  on division of 30 by 6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(30)
		calculator.Divide(6)
		assert.Equal(t, 5.00, calculator.currentValue)
	})

	t.Run("should return -5.00  on division of 30 by -6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(30)
		calculator.Divide(-6)
		assert.Equal(t, -5.00, calculator.currentValue)
	})

	t.Run("should return -5.00  on division of -30 by 6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-30)
		calculator.Divide(6)
		assert.Equal(t, -5.00, calculator.currentValue)
	})

	t.Run("should return 5.00  on division of -30 by -6", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(-30)
		calculator.Divide(-6)
		assert.Equal(t, 5.00, calculator.currentValue)
	})

	t.Run("should return 5.15 on division of 25.75 by 5", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.Add(25.75)
		calculator.Divide(5)
		assert.Equal(t, 5.15, calculator.currentValue)
	})
}
