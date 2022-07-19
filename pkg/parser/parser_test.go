package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("should return string and float values", func(t *testing.T) {
		operation, value := Parser("")
		assert.IsType(t, "", operation)
		assert.IsType(t, float64(0), value)
	})
}
