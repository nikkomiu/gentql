package cmd

import (
	"bytes"
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func executeWithArgs(ctx context.Context, args []string) (stdout string, stderr string, err error) {
	var outBuf bytes.Buffer
	var errBuf bytes.Buffer

	err = Execute(ctx, WithOutput(&outBuf, &errBuf), WithArgs(args))

	return outBuf.String(), errBuf.String(), err
}

func ContextT(t *testing.T) context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)
	return ctx
}

func TestRootCmd(t *testing.T) {
	// Arrange
	ctx := ContextT(t)

	// Act
	stdout, stderr, err := executeWithArgs(ctx, []string{})

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, stdout, "GentQL backend application services.")
	assert.Contains(t, stdout, "Usage:")
	assert.Equal(t, stderr, "")
}
