package feature

import (
	"context"
	"net"

	"github.com/Microsoft/go-winio"
)

func path() string {
	return `\\.\pipe\dockerBackendApiServer`
}

func dial(ctx context.Context) (net.Conn, error) {
	return winio.DialPipeContext(ctx, path())
}
