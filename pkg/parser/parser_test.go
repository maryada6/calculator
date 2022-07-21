package parser

import (
	"calculator/pkg/handler"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParser(t *testing.T) {
	t.Run("should return tokens as string array for given input", func(t *testing.T) {
		tokens := Split("add 5")
		assert.Equal(t, "add", tokens[0])
		assert.Equal(t, "5", tokens[1])
		assert.Equal(t, 2, len(tokens))

		tokensTwo := Split("subtract 4")
		assert.Equal(t, "subtract", tokensTwo[0])
		assert.Equal(t, "4", tokensTwo[1])
		assert.Equal(t, 2, len(tokensTwo))

		tokensThree := Split("this is random 56")
		assert.Equal(t, "this", tokensThree[0])
		assert.Equal(t, "is", tokensThree[1])
		assert.Equal(t, "random", tokensThree[2])
		assert.Equal(t, "56", tokensThree[3])
		assert.Equal(t, 4, len(tokensThree))
	})
}

func TestValidator(t *testing.T) {
	t.Run("should return false is tokens  have more than 2 tokens", func(t *testing.T) {
		tokens := []string{"add", "5", "up"}
		assert.Equal(t, false, Validator(tokens))
	})

	t.Run("should return true for valid tokens", func(t *testing.T) {
		tokensOne := []string{"add", "5"}
		assert.Equal(t, true, Validator(tokensOne))
		tokensTwo := []string{"addition", "5"}
		assert.Equal(t, true, Validator(tokensTwo))
	})
}

func TestParseInput(t *testing.T) {
	t.Run("should return add and 5 for add 5", func(t *testing.T) {
		parser := ParseInput("add 5")
		assert.Equal(t, handler.Operation("add"), parser.Operation)
		assert.Equal(t, 5.0, parser.Value)
	})

	t.Run("should return cancel and 0 for cancel", func(t *testing.T) {
		parser := ParseInput("cancel")
		assert.Equal(t, handler.Operation("cancel"), parser.Operation)
		assert.Equal(t, 0.0, parser.Value)
	})

	t.Run("should return subtract and 0 for subtract", func(t *testing.T) {
		parser := ParseInput("cancel")
		assert.Equal(t, handler.Operation("cancel"), parser.Operation)
		assert.Equal(t, 0.0, parser.Value)
	})

	t.Run("should return empty string and 0.0 for invalid input", func(t *testing.T) {
		parser := ParseInput("ok this is me")
		assert.Equal(t, handler.Operation(""), parser.Operation)
		assert.Equal(t, 0.0, parser.Value)
	})
}
