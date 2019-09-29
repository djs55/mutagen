package filesystem

import (
	"fmt"
	"os"
	"path/filepath"
)

func mutagenData(pathComponents ...string) (string, error) {
	localAppData := os.Getenv("LOCALAPPDATA")
	if localAppData == "" {
		return "", fmt.Errorf("unable to get 'LOCALAPPDATA'")
	}

	dir := filepath.Join(localAppData, "Docker", "mutagen")
	if fi, err := os.Stat(dir); err == nil && fi.IsDir() {
		return dir, nil
	}
	return "", fmt.Errorf("%s does not exist", dir)
}
