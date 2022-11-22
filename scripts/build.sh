#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# CoinflectChain root folder
COINFLECT_PATH=$( cd "$( dirname "${BASH_SOURCE[0]}" )"; cd .. && pwd )
# Load the versions
source "$COINFLECT_PATH"/scripts/versions.sh
# Load the constants
source "$COINFLECT_PATH"/scripts/constants.sh

# Download dependencies
echo "Downloading dependencies..."
go mod download

# Build coinflectchain
"$COINFLECT_PATH"/scripts/build_coinflect.sh

# Build coreth
"$COINFLECT_PATH"/scripts/build_coreth.sh

# Exit build successfully if the binaries are created
if [[ -f "$coinflectchain_path" && -f "$evm_path" ]]; then
        echo "Build Successful"
        exit 0
else
        echo "Build failure" >&2
        exit 1
fi
