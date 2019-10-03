package desktop

import (
	"testing"

	urlpkg "github.com/mutagen-io/mutagen/pkg/url"
	"github.com/stretchr/testify/assert"
)

func TestMacPath(t *testing.T) {
	p := pathInVM(&urlpkg.URL{
		Path: `/Users/foo`,
	})
	assert.Equal(t, "/mutagen/Users/foo", p)
}
