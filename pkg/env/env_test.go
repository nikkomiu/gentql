package env_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nikkomiu/gentql/pkg/env"
)

func TestStr(t *testing.T) {
	key := "TEST_ENV_STR"
	defaultValue := "DEFAULT_STRING_VALUE"

	tt := []struct {
		name     string
		value    string
		expected string
	}{
		{name: "base", value: "some", expected: "some"},
		{name: "default", expected: defaultValue},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			if tc.value != "" {
				t.Setenv(key, tc.value)
			}

			// Act
			val := env.Str(key, defaultValue)

			// Assert
			assert.Equal(t, tc.expected, val)
		})
	}
}

func TestInt(t *testing.T) {
	key := "TEST_ENV_INT"
	defaultValue := 33

	tt := []struct {
		name     string
		value    string
		expected int
	}{
		{name: "base", value: "30", expected: 30},
		{name: "bad value", value: "string", expected: defaultValue},
		{name: "default", expected: defaultValue},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			if tc.value != "" {
				t.Setenv(key, tc.value)
			}

			// Act
			val := env.Int(key, defaultValue)

			// Assert
			assert.Equal(t, tc.expected, val)
		})
	}
}

func TestDuration(t *testing.T) {
	key := "TEST_ENV_DURATION"
	defaultValue := time.Minute

	tt := []struct {
		name     string
		value    string
		expected time.Duration
	}{
		{name: "base", value: "1m30s", expected: time.Minute + (30 * time.Second)},
		{name: "bad value", value: "string", expected: defaultValue},
		{name: "default", expected: defaultValue},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			// Arrange
			if tc.value != "" {
				t.Setenv(key, tc.value)
			}

			// Act
			val := env.Duration(key, defaultValue)

			// Assert
			assert.Equal(t, tc.expected, val)
		})
	}
}
