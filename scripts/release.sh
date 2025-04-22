#!/bin/bash

# Helper script to create a new release

set -e

if [ $# -ne 1 ]; then
  echo "Usage: $0 <version>"
  echo "Example: $0 v1.0.0"
  exit 1
fi

VERSION=$1

# Validate version format (vX.Y.Z)
if ! [[ $VERSION =~ ^v[0-9]+\.[0-9]+\.[0-9]+$ ]]; then
  echo "Error: Version must be in format vX.Y.Z (e.g., v1.0.0)"
  exit 1
fi

# Check if tag already exists
if git rev-parse "$VERSION" >/dev/null 2>&1; then
  echo "Error: Tag $VERSION already exists"
  exit 1
fi

# Run tests
echo "Running tests..."
go test -v ./...

# Create and push tag
echo "Creating tag $VERSION..."
git tag -a "$VERSION" -m "Release $VERSION"
git push origin "$VERSION"

echo "Release $VERSION created and pushed to origin."
echo "Don't forget to create a release on GitHub!"
echo "https://github.com/vaidehimani/gotextrim/releases/new?tag=$VERSION"