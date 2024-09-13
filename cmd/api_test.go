package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAPICmd(t *testing.T) {
	t.Setenv("PORT", "9901")

	tt := []struct {
		name     string
		dbDriver string

		expectedErr    string
		expectedStderr string
	}{
		{
			name: "base",
		},
		{
			name:     "bad database driver",
			dbDriver: "bad_driver",

			expectedErr:    "unsupported driver: \"bad_driver\"",
			expectedStderr: "Error: unsupported driver: \"bad_driver\"\n",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			if tc.dbDriver != "" {
				t.Setenv("DATABASE_DRIVER", tc.dbDriver)
			}
			ctx, cancel := context.WithTimeout(ContextT(t), 250*time.Millisecond)
			defer cancel()

			// Act
			stdout, stderr, err := executeWithArgs(ctx, []string{"api"})

			// Assert
			if tc.expectedErr == "" {
				assert.NoError(t, err)
			} else {
				assert.EqualError(t, err, tc.expectedErr)
			}
			assert.Equal(t, "", stdout)
			assert.Equal(t, tc.expectedStderr, stderr)
		})
	}

	t.Run("routes", testAPICmdRoutes)
}

func testAPICmdRoutes(t *testing.T) {
	go func() {
		ctx := ContextT(t)
		_, _, err := executeWithArgs(ctx, []string{"api"})
		t.Logf("api command: err:%s\n", err)
	}()

	buf := bytes.NewBufferString("{\"query\": \"query{ping}\"}")
	resp, err := http.Post("http://localhost:9901/graphql", "application/json", buf)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	defer resp.Body.Close()

	var respData map[string]any
	json.NewDecoder(resp.Body).Decode(&respData)
	assert.NoError(t, err)
	assert.Equal(t, map[string]any{"data": map[string]any{"ping": "pong"}}, respData)
}
