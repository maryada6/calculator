package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReader(t *testing.T) {
	t.Run("should return a string", func(t *testing.T) {
		assert.IsType(t, "", Reader())
	})
}
