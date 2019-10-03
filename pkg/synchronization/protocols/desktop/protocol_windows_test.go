package desktop

import (
	"testing"

	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
	"github.com/stretchr/testify/assert"
)

func TestWindowsPath(t *testing.T) {
	p := pathInVM(&urlpkg.URL{
		Path: `C:\Users`,
	})
	assert.Equal(t, "/mutagen/C/Users", p)
}
