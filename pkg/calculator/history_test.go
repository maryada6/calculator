package calculator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStoreInput(t *testing.T) {
	t.Run("should append the input", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.SaveInput("add", 5)
		calculator.SaveInput("subtract", 5)
		assert.Equal(t, 2, len(history))
	})

	t.Run("should append the input", func(t *testing.T) {
		calculator := NewCalculator()
		calculator.SaveInput("multiply", 5)
		calculator.SaveInput("subtract", 5)
		assert.Equal(t, 2, len(history))
	})
}
