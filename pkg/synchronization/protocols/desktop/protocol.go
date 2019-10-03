package desktop

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"path"

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
	conn, err := Dial()
	if err != nil {
		return nil, err
	}
	hash := md5.New()
	dir := "desktop" + hex.EncodeToString(hash.Sum([]byte(url.Path)))
	// Create the endpoint client.
	return remote.NewEndpointClient(conn, path.Join("/mutagen", dir), session, version, configuration, alpha)
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
