#!/bin/sh

set -e

if [ ! -f "build/env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
workspace="$PWD/build/_workspace"
root="$PWD"
kcoindir="$workspace/src/github.com/kowala-tech/kcoin"
if [ ! -L "$kcoindir/client" ]; then
    mkdir -p "$kcoindir"
    cd "$kcoindir"
    ln -s ../../../../../../. client
    cd "$root"
fi

# Set up the environment to use the workspace.
GOPATH="$workspace"
export GOPATH

# Run the command inside the workspace.
cd "$kcoindir/client"
PWD="$kcoindir/client"

# Launch the arguments with the configured environment.
exec "$@"
