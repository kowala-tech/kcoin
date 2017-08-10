#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
ethdir="$workspace/src/github.com/kowala-tech"
if [ ! -L "$ethdir/kUSD" ]; then
    mkdir -p "$ethdir"
    cd "$ethdir"
    ln -s ../../../../../. kUSD
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$ethdir/kUSD"
PWD="$ethdir/kUSD"

# Launch the arguments with the configured environment.
exec "$@"
