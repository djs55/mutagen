package desktop

import (
	"encoding/binary"
	"encoding/hex"
	"os"
	"path"
	"path/filepath"
	"strings"
	"unicode"

	"github.com/mutagen-io/mutagen/pkg/logging"
	"github.com/mutagen-io/mutagen/pkg/synchronization"
	"github.com/mutagen-io/mutagen/pkg/synchronization/endpoint/remote"
	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
)

const mutagenDirInVM = "/mutagen"

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
	// Create the endpoint client.
	return remote.NewEndpointClient(conn, pathInVM(url), session, version, configuration, alpha)
}

func pathInVM(url *urlpkg.URL) string {
	p := url.Path
	if len(p) >= 2 && unicode.IsLetter(rune(p[0])) && string(p[1]) == ":" {
		// starts with a Windows drive letter, replace with a path separator
		p = filepath.Join(string(p[0]), p[2:])
	}
	elements := []string{mutagenDirInVM}
	for _, element := range strings.Split(p, string(os.PathSeparator)) {
		elements = append(elements, element)
	}
	// convert to Go/Unix style path
	return path.Join(elements...)
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
