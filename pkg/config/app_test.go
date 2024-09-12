package config_test

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nikkomiu/gentql/pkg/config"
)

func unsetenv(t *testing.T, key string) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return
	}

	if err := os.Unsetenv(key); err != nil {
		t.Fatalf("failed to unset env var (%s): %s", key, err)
	}
	t.Cleanup(func() {
		os.Setenv(key, val)
	})
}

func TestApp(t *testing.T) {
	t.Run("initial default", testAppInitialDefault)
	t.Run("initial override", testAppInitialOverride)
	t.Run("existing", testAppExisting)
}

func testAppInitialDefault(t *testing.T) {
	// Arrange
	unsetenv(t, "ADDRESS")
	unsetenv(t, "PORT")
	unsetenv(t, "SERVER_SHUTDOWN_TIMEOUT")
	unsetenv(t, "DATABASE_DRIVER")
	unsetenv(t, "DATABASE_URL")
	ctx, _ := config.WithApp(context.Background())

	// Act
	cfg := config.AppFromContext(ctx)

	// Assert
	assert.Equal(t, cfg.Server.Host, "")
	assert.Equal(t, cfg.Server.Port, 8080)
	assert.Equal(t, cfg.Server.Addr(), ":8080")
	assert.Equal(t, cfg.Server.DisplayAddr(), "http://localhost:8080/")
	assert.Equal(t, cfg.Server.ShutdownTimeout, 10*time.Second)
	assert.Equal(t, cfg.Database.Driver, "postgres")
	assert.Equal(t, cfg.Database.URL, "postgres://localhost/gentql_dev?sslmode=disable")
}

func testAppInitialOverride(t *testing.T) {
	// Arrange
	t.Setenv("ADDRESS", "10.0.0.10")
	t.Setenv("PORT", "9999")
	t.Setenv("SERVER_SHUTDOWN_TIMEOUT", "1m")
	t.Setenv("DATABASE_DRIVER", "sqlite")
	t.Setenv("DATABASE_URL", "file:ent?mode=memory&_fk=1")
	ctx, _ := config.WithApp(context.Background())

	// Act
	cfg := config.AppFromContext(ctx)

	// Assert
	assert.Equal(t, cfg.Server.Host, "10.0.0.10")
	assert.Equal(t, cfg.Server.Port, 9999)
	assert.Equal(t, cfg.Server.Addr(), "10.0.0.10:9999")
	assert.Equal(t, cfg.Server.DisplayAddr(), "http://10.0.0.10:9999/")
	assert.Equal(t, cfg.Server.ShutdownTimeout, time.Minute)
	assert.Equal(t, cfg.Database.Driver, "sqlite")
	assert.Equal(t, cfg.Database.URL, "file:ent?mode=memory&_fk=1")
}

func testAppExisting(t *testing.T) {
	// Arrange
	unsetenv(t, "ADDRESS")
	unsetenv(t, "PORT")
	unsetenv(t, "SERVER_SHUTDOWN_TIMEOUT")
	unsetenv(t, "DATABASE_DRIVER")
	unsetenv(t, "DATABASE_URL")
	ctx, _ := config.WithApp(context.Background())
	t.Setenv("ADDRESS", "10.0.0.11")

	// Act
	cfg := config.AppFromContext(ctx)

	// Assert
	assert.Equal(t, cfg.Server.Host, "")
}
