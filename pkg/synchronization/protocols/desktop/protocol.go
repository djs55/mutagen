package desktop

import (
	"context"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"strings"

	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	"github.com/mutagen-io/mutagen/pkg/synchronization/endpoint/remote"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

// protocolHandler implements the session.ProtocolHandler interface for
// connecting to remote endpoints over SSH. It uses the agent infrastructure
// over an SSH transport.
type protocolHandler struct{}

// Dial connects to an SSH endpoint.
func (h *protocolHandler) Connect(
	logger *logging.Logger,
	url *urlpkg.URL,
	prompter string,
	session string,
	version synchronization.Version,
	configuration *synchronization.Configuration,
	alpha bool,
) (synchronization.Endpoint, error) {

	d := net.Dialer{}

	conn, err := d.DialContext(context.Background(), "unix", os.ExpandEnv("$HOME/Library/Containers/com.docker.docker/Data/vms/0/connect"))
	if err != nil {
		return nil, err
	}

	// request the setup service
	remote1 := strings.NewReader(fmt.Sprintf("00000003.%08x\n", 4099))
	_, err = io.Copy(conn, remote1)
	if err != nil {
		return nil, err
	}

	// Create the endpoint client.
	return remote.NewEndpointClient(conn, filepath.Join("/mutagen", url.Path), session, version, configuration, alpha)
}

func toHexBigE(v uint32) string {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, v)
	return hex.EncodeToString(bytes)
}

func init() {
	// Register the SSH protocol handler with the synchronization package.
	synchronization.ProtocolHandlers[urlpkg.Protocol_Desktop] = &protocolHandler{}
}
