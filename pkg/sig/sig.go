package sig

import (
	"context"
	"net/http"
	"time"
)

func ListenAndServe(ctx context.Context, addr string, handler http.Handler, shutdownTimeout time.Duration) error {
	server := &http.Server{Addr: addr, Handler: handler}

	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err

	case <-ctx.Done():
		// do nothing
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	return server.Shutdown(ctx)
}
