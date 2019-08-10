package desktop

import (
	"testing"

	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
	"github.com/stretchr/testify/assert"
)

func TestWindowsPath(t *testing.T) {
	u := &urlpkg.URL{
		Path: `C:\desktopHASH`,
	}
	p := pathInVM(u)
	assert.Equal(t, "/mutagen/desktopHASH", p)
}

func TestMacPath(t *testing.T) {
	u := &urlpkg.URL{
		Path: `desktopHASH`,
	}
	p := pathInVM(u)
	assert.Equal(t, "/mutagen/desktopHASH", p)
}
