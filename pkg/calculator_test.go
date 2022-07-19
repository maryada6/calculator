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
}
