package filesystem

import (
	"io/ioutil"
	"os"
	"runtime"
	"testing"
	"unicode/utf8"
)

// TestPathSeparatorSingleByte verifies that the platform path separator rune is
// encoded as a single byte in UTF-8. We rely on this assumption for high
// performance in ensureValidName.
func TestPathSeparatorSingleByte(t *testing.T) {
	if utf8.RuneLen(os.PathSeparator) != 1 {
		t.Fatal("OS path separator does not have single-byte UTF-8 encoding")
	}
}

func TestDirectoryContentsNotExist(t *testing.T) {
	if _, err := DirectoryContentsByPath("/does/not/exist"); err == nil {
		t.Error("directory listing succeedeed for non-existent path")
	}
}

func TestDirectoryContentsFile(t *testing.T) {
	// Create an empty temporary file and defer its cleanup.
	file, err := ioutil.TempFile("", "mutagen_filesystem")
	if err != nil {
		t.Fatal("unable to create temporary file:", err)
	} else if err = file.Close(); err != nil {
		t.Error("unable to close temporary file:", err)
	}
	defer os.Remove(file.Name())

	// Ensure that directory listing fails.
	if _, err := DirectoryContentsByPath(file.Name()); err == nil {
		t.Error("directory listing succeedeed for non-directory path")
	}
}

func TestDirectoryContentsGOROOT(t *testing.T) {
	if contents, err := DirectoryContentsByPath(runtime.GOROOT()); err != nil {
		t.Fatal("directory listing failed for GOROOT:", err)
	} else if contents == nil {
		t.Fatal("directory contents nil for GOROOT")
	}
}

// TestDirectorySymbolicLinkRemoval tests that removal of symbolic links that
// point to directories works as expected.
func TestDirectorySymbolicLinkRemoval(t *testing.T) {
	// Create a temporary directory and defer its cleanup.
	temporaryDirectoryPath, err := ioutil.TempDir("", "mutagen_filesystem_test")
	if err != nil {
		t.Fatal("unable to create temporary directory:", err)
	}
	defer os.RemoveAll(temporaryDirectoryPath)

	// Create a directory handle for this temporary directory and defer its
	// closure.
	directory, _, err := OpenDirectory(temporaryDirectoryPath, false)
	if err != nil {
		t.Fatal("unable to open directory handle:", err)
	}
	defer directory.Close()

	// Create a directory that will serve as our target.
	if err := directory.CreateDirectory("target"); err != nil {
		t.Fatal("unable to create target directory:", err)
	}

	// Open the target directory and defer its closure.
	target, err := directory.OpenDirectory("target")
	if err != nil {
		t.Fatal("unable to open target directory:", err)
	}
	defer target.Close()

	// Create content within the target.
	if err := target.CreateDirectory("content"); err != nil {
		t.Fatal("unable to create content in target directory:", err)
	}

	// Create a symbolic link to the target.
	if err := directory.CreateSymbolicLink("link", "target"); err != nil {
		t.Fatal("unable to create symbolic link:", err)
	}

	// Remove the symbolic link.
	if err := directory.RemoveSymbolicLink("link"); err != nil {
		t.Fatal("unable to remove symbolic link:", err)
	}

	// Grab the target metadata to ensure that it still exists.
	if _, err := directory.ReadContentMetadata("target"); err != nil {
		t.Error("unable to read target metadata:", err)
	}
}
