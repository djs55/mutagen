// +build !windows

package filesystem

import (
	"os"
	"path/filepath"
)

func mutagenData(pathComponents ...string) (string, error) {
	return filepath.Join(os.ExpandEnv("$HOME/Library/Group Containers/group.com.docker/mutagen"), filepath.Join(pathComponents...)), nil
}
