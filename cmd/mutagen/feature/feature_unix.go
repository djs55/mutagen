// +build !windows

package feature

import (
	"context"
	"net"
	"os"
)

func path() string {
	return os.ExpandEnv("$HOME/Library/Containers/com.docker.docker/Data/docker-api.sock")
}

func dial(ctx context.Context) (net.Conn, error) {
	dialer := &net.Dialer{}
	return dialer.DialContext(ctx, "unix", path())
}
