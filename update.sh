#!/bin/sh
set -e

# Build mutagen binaries, upload to S3

VERSION=$(git rev-parse HEAD)

PUTME=~/bin/putme

if [ ! -x "$PUTME" ]; then
  echo "Please set PUTME to the path to a compiled putme binary"
  exit 1
fi

rm -rf build
go run scripts/build.go --mode=release
./build/mutagen mutagen legal > build/OSS-LICENSES

${PUTME} --component mutagen --os mac --version "$VERSION" --key "docker-mutagen" --path ./build/mutagen
${PUTME} --component mutagen --os mac --version "$VERSION" --key "docker-mutagen.exe" --path ./build/cli/windows_amd64
${PUTME} --component mutagen --os mac --version "$VERSION" --key "mutagen-agent" --path ./build/agent/linux_amd64
${PUTME} --component mutagen --os mac --version "$VERSION" --key "OSS-LICENSES" --path ./build/OSS-LICENSES
