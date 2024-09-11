package env

import (
	"os"
	"strconv"
)

func Str(key, defaultValue string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultValue
}

func Int(key string, defaultValue int) int {
	if val, err := strconv.Atoi(os.Getenv(key)); err == nil {
		return val
	}

	return defaultValue
}
