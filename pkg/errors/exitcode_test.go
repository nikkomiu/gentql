package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nikkomiu/gentql/pkg/errors"
)

func TestExitCode(t *testing.T) {
	tt := []struct {
		name     string
		exitCode int
		innerErr error
	}{
		{
			name:     "default",
			exitCode: 51,
			innerErr: fmt.Errorf("simple error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Act
			err := errors.NewExitCode(tc.innerErr, tc.exitCode)

			// Assert
			assert.Equal(t, tc.exitCode, err.ExitCode())
			assert.Equal(t, tc.innerErr, err.Unwrap())
			assert.Equal(t, tc.innerErr.Error(), err.Error())
			assert.Equal(t, tc.innerErr.Error(), err.String())
		})
	}
}
