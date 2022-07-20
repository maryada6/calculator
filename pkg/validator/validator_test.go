package validator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStringInList(t *testing.T) {
	t.Run("should return true for a string in our given list", func(t *testing.T) {
		var commandList = []string{"add", "subtract", "multiply", "divide", "cancel", "exit"}
		assert.Equal(t, true, stringInList("add", commandList))
	})

	t.Run("should return false for a string not in our given list", func(t *testing.T) {
		var commandList = []string{"add", "subtract", "multiply", "divide", "cancel", "exit"}
		assert.Equal(t, false, stringInList("random", commandList))
	})
}
