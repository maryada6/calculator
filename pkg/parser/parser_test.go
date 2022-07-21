package parser

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("should return tokens as string array for given input", func(t *testing.T) {
		tokens := Parser("add 5")
		assert.Equal(t, "add", tokens[0])
		assert.Equal(t, "5", tokens[1])
		assert.Equal(t, 2, len(tokens))

		tokensTwo := Parser("subtract 4")
		assert.Equal(t, "subtract", tokensTwo[0])
		assert.Equal(t, "4", tokensTwo[1])
		assert.Equal(t, 2, len(tokensTwo))

		tokensThree := Parser("this is random 56")
		assert.Equal(t, "this", tokensThree[0])
		assert.Equal(t, "is", tokensThree[1])
		assert.Equal(t, "random", tokensThree[2])
		assert.Equal(t, "56", tokensThree[3])
		assert.Equal(t, 4, len(tokensThree))
	})
}

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

	t.Run("should return false for invalid tokens", func(t *testing.T) {
		tokens := []string{"it", "with", "5"}
		assert.Equal(t, false, Validator(tokens))
	})
}

func TestParseInput(t *testing.T) {
	t.Run("should return add and 5 for add 5", func(t *testing.T) {
		operation, value := ParseInput("add 5")
		assert.Equal(t, "add", operation)
		assert.Equal(t, 5.0, value)
	})

	t.Run("should return cancel and 0 for cancel", func(t *testing.T) {
		operation, value := ParseInput("cancel")
		assert.Equal(t, "cancel", operation)
		assert.Equal(t, 0.0, value)
	})

	t.Run("should return subtract and 0 for subtract", func(t *testing.T) {
		operation, value := ParseInput("cancel")
		assert.Equal(t, "cancel", operation)
		assert.Equal(t, 0.0, value)
	})

	t.Run("should return error for invalid input", func(t *testing.T) {
		assert.NotPanics(t, func() {
			ParseInput("this is rando string")
		})
	})
}
