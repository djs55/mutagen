package filesystem

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pkg/errors"
)

func mutagenData(pathComponents ...string) (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return "", fmt.Errorf("unable to get 'LOCALAPPDATA'")
	}

	dir := filepath.Join(filepath.Join(localAppData, "Docker", "mutagen"), filepath.Join(pathComponents...))
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", errors.Wrapf(err, "unable to create %s", dir)
	}

	return dir, nil
}
