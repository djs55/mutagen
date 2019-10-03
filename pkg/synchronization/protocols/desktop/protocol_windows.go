package desktop

import (
	"net"
	"time"

	"github.com/Microsoft/go-winio"
)

const pipeName = `\\.\pipe\dockerMutagen`

// Dial the mutagen server.
func Dial() (net.Conn, error) {
	timeout := 1 * time.Second
	return winio.DialPipe(pipeName, &timeout)
}
