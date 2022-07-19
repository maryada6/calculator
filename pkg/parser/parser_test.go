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
