package reader

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestReader(t *testing.T) {
	t.Run("should return the string read by the reader", func(t *testing.T) {
		reader := strings.NewReader("add 5")
		assert.Equal(t, "add 5", Reader(reader))
		readerTwo := strings.NewReader("subtract 5")
		assert.Equal(t, "subtract 5", Reader(readerTwo))
		readerThree := strings.NewReader("multiply 4")
		assert.Equal(t, "multiply 4", Reader(readerThree))
		readerFour := strings.NewReader("something random")
		assert.Equal(t, "something random", Reader(readerFour))
	})
}
