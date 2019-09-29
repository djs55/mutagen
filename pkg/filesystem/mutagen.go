package filesystem

import (
	"os"

	"github.com/pkg/errors"
)

const (
	// MutagenDataDirectoryName is the name of the global Mutagen data directory
	// inside the user's home directory.
	MutagenDataDirectoryName = ".mutagen"

	// MutagenDataDirectoryDevelopmentName is the name of the global Mutagen
	// data directory inside the user's home directory for development builds of
	// Mutagen.
	MutagenDataDirectoryDevelopmentName = ".mutagen-dev"

	// MutagenLegacyGlobalSynchronizationConfigurationName is the name of the
	// legacy global Mutagen synchronization configuration file inside the
	// user's home directory.
	MutagenLegacyGlobalSynchronizationConfigurationName = ".mutagen.toml"

	// MutagenGlobalConfigurationName is the name of the global Mutagen
	// configuration file inside the user's home directory.
	MutagenGlobalConfigurationName = ".mutagen.yml"

	// MutagenDaemonDirectoryName is the name of the daemon storage directory
	// within the Mutagen data directory.
	MutagenDaemonDirectoryName = "daemon"

	// MutagenAgentsDirectoryName is the name of the agent storage directory
	// within the Mutagen data directory.
	MutagenAgentsDirectoryName = "agents"

	// MutagenSynchronizationSessionsDirectoryName is the name of the
	// synchronization session storage directory within the Mutagen data
	// directory.
	MutagenSynchronizationSessionsDirectoryName = "sessions"

	// MutagenSynchronizationCachesDirectoryName is the name of the
	// synchronization cache storage directory within the Mutagen data
	// directory.
	MutagenSynchronizationCachesDirectoryName = "caches"

	// MutagenSynchronizationArchivesDirectoryName is the name of the
	// synchronization archive storage directory within the Mutagen data
	// directory.
	MutagenSynchronizationArchivesDirectoryName = "archives"

	// MutagenSynchronizationStagingDirectoryName is the name of the
	// synchronization staging storage directory within the Mutagen data
	// directory.
	MutagenSynchronizationStagingDirectoryName = "staging"

	// MutagenForwardingDirectoryName is the name of the forwarding data
	// directory within the Mutagen data directory.
	MutagenForwardingDirectoryName = "forwarding"
)

// Mutagen computes (and optionally creates) subdirectories inside the Mutagen
// data directory.
func Mutagen(create bool, pathComponents ...string) (string, error) {

	// Compute the target path.
	result, err := mutagenData(pathComponents...)
	if err != nil {
		return "", err
	}

	// Handle directory creation, if requested.
	//
	// TODO: Should we iterate through each component and ensure the user hasn't
	// changed the directory permissions? MkdirAll won't reset them. But I
	// suppose the user may have changed them for whatever reason (though I
	// can't imagine any).
	if create {
		// Create the directory.
		if err := os.MkdirAll(result, 0700); err != nil {
			return "", errors.Wrap(err, "unable to create subpath")
		}
	}

	// Success.
	return result, nil
}
