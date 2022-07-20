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

func TestValidator(t *testing.T) {
	t.Run("should return false is tokens does not have 2 tokens", func(t *testing.T) {
		tokens := []string{"add", "5", "up"}
		assert.Equal(t, false, Validator(tokens))
	})

	t.Run("should return true for valid tokens", func(t *testing.T) {
		tokens := []string{"add", "5"}
		assert.Equal(t, true, Validator(tokens))
	})

	t.Run("should return false for invalid tokens", func(t *testing.T) {
		tokens := []string{"added", "5"}
		assert.Equal(t, false, Validator(tokens))
	})
}
