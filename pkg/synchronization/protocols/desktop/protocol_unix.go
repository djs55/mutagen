// +build !windows

package desktop

import (
	"context"
	"net"
	"os"
)

// Dial the mutagen server.
func Dial() (net.Conn, error) {
	d := net.Dialer{}

	return d.DialContext(context.Background(), "unix", os.ExpandEnv("$HOME/Library/Containers/com.docker.docker/Data/mutagen.sock"))
}
