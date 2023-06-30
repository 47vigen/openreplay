#!/bin/bash

# Set the remote upstream URL
upstream_url="https://github.com/47vigen/openreplay"

# Check if the remote upstream exists
if ! git ls-remote --exit-code --quiet upstream &> /dev/null; then
  echo "Remote upstream not found. Adding remote upstream..."
  git remote add upstream "$upstream_url"
  echo "Remote upstream added."
fi

# Perform git pull from upstream/main
echo "Pulling changes from upstream/main..."
git pull upstream main

# Perform git pull from upstream/main
echo "Pulling changes from upstream/main..."
git pull upstream main

echo "Done."
