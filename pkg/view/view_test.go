package view

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"os"
	"testing"
)

func TestInitialSign(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	InitialSign()
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, ">", string(bytes))
}

func TestView(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	View(5.00)
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "5\n", string(bytes))
}

func TestInvalidInput(t *testing.T) {
	defer setupStdout()()
	r, fakeStdout, err := os.Pipe()
	require.NoError(t, err)
	os.Stdout = fakeStdout
	InvalidInput("Invalid Input")
	fakeStdout.Close()
	bytes, err := io.ReadAll(r)
	r.Close()
	assert.Equal(t, "Invalid Input\n", string(bytes))
}

func setupStdout() func() {
	originalStdout := os.Stdout
	return func() {
		os.Stdout = originalStdout
	}
}
