#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
kusddir="$workspace/src/github.com/kowala-tech"
if [ ! -L "$kusddir/kUSD" ]; then
    mkdir -p "$kusddir"
    cd "$kusddir"
    ln -s ../../../../../. kUSD
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$kusddir/kUSD"
PWD="$kusddir/kUSD"

# Launch the arguments with the configured environment.
exec "$@"
