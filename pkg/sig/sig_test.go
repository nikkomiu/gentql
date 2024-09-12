package sig_test

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/nikkomiu/gentql/pkg/sig"
)

func TestListenAndServe(t *testing.T) {
	tt := []struct {
		name    string
		addr    string
		handler http.Handler

		shutdownTimeout time.Duration

		wantErr bool
	}{
		{
			name: "base",
			addr: ":9990",
		},
		{
			name: "start error",
			addr: ":no_port",

			wantErr: true,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			err := sig.ListenAndServe(ctx, tc.addr, tc.handler, tc.shutdownTimeout)

			assert.Equal(t, tc.wantErr, err != nil, err)
		})
	}
}
