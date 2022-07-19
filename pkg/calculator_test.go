package pkg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewCalculator(t *testing.T) {
	t.Run("should initialise the calculator entity with value 0", func(t *testing.T) {
		assert.NotPanics(t, func() {
			NewCalculator(0)
		})
	})

	t.Run("Should return calculator entity", func(t *testing.T) {
		assert.IsType(t, Calculator{}, NewCalculator(0))
	})
}
