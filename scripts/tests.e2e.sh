#!/usr/bin/env bash
set -e
set -o nounset
set -o pipefail

# e.g.,
# ./scripts/build.sh
# ./scripts/tests.e2e.sh ./build/coinflectchain
# ENABLE_WHITELIST_VTX_TESTS=true ./scripts/tests.e2e.sh ./build/coinflectchain
if ! [[ "$0" =~ scripts/tests.e2e.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

COINFLECTCHAIN_PATH="${1-}"
if [[ -z "${COINFLECTCHAIN_PATH}" ]]; then
  echo "Missing COINFLECTCHAIN_PATH argument!"
  echo "Usage: ${0} [COINFLECTCHAIN_PATH]" >> /dev/stderr
  exit 255
fi

# Set the CGO flags to use the portable version of BLST
#
# We use "export" here instead of just setting a bash variable because we need
# to pass this flag to all child processes spawned by the shell.
export CGO_CFLAGS="-O -D__BLST_PORTABLE__"

ENABLE_WHITELIST_VTX_TESTS=${ENABLE_WHITELIST_VTX_TESTS:-false}
# ref. https://onsi.github.io/ginkgo/#spec-labels
GINKGO_LABEL_FILTER="!whitelist-tx"
if [[ ${ENABLE_WHITELIST_VTX_TESTS} == true ]]; then
  # run only "whitelist-tx" tests, no other test
  GINKGO_LABEL_FILTER="whitelist-tx"
fi
echo GINKGO_LABEL_FILTER: ${GINKGO_LABEL_FILTER}

#################################
# download coinflect-network-runner
# https://github.com/coinflect/coinflect-network-runner
# TODO: migrate to upstream coinflect-network-runner
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
NETWORK_RUNNER_VERSION=1.3.1
DOWNLOAD_PATH=/tmp/coinflect-network-runner.tar.gz
DOWNLOAD_URL="https://github.com/coinflect/coinflect-network-runner/releases/download/v${NETWORK_RUNNER_VERSION}/coinflect-network-runner_${NETWORK_RUNNER_VERSION}_${GOOS}_${GOARCH}.tar.gz"

rm -f ${DOWNLOAD_PATH}
rm -f /tmp/coinflect-network-runner

echo "downloading coinflect-network-runner ${NETWORK_RUNNER_VERSION} at ${DOWNLOAD_URL}"
curl --fail -L ${DOWNLOAD_URL} -o ${DOWNLOAD_PATH}

echo "extracting downloaded coinflect-network-runner"
tar xzvf ${DOWNLOAD_PATH} -C /tmp
/tmp/coinflect-network-runner -h

GOPATH="$(go env GOPATH)"
PATH="${GOPATH}/bin:${PATH}"

#################################
echo "building e2e.test"
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@v2.1.4
ACK_GINKGO_RC=true ginkgo build ./tests/e2e
./tests/e2e/e2e.test --help

#################################
# run "coinflect-network-runner" server
echo "launch coinflect-network-runner in the background"
/tmp/coinflect-network-runner \
server \
--log-level debug \
--port=":12342" \
--disable-grpc-gateway 2> /dev/null &
PID=${!}

#################################
echo "running e2e tests against the local cluster with ${COINFLECTCHAIN_PATH}"
./tests/e2e/e2e.test \
--ginkgo.v \
--log-level debug \
--network-runner-grpc-endpoint="0.0.0.0:12342" \
--network-runner-coinflectchain-path=${COINFLECTCHAIN_PATH} \
--network-runner-coinflectchain-log-level="WARN" \
--test-keys-file=tests/test.insecure.secp256k1.keys --ginkgo.label-filter="${GINKGO_LABEL_FILTER}" \
&& EXIT_CODE=$? || EXIT_CODE=$?

kill ${PID}

if [[ ${EXIT_CODE} -gt 0 ]]; then
  echo "FAILURE with exit code ${EXIT_CODE}"
  exit ${EXIT_CODE}
else
  echo "ALL SUCCESS!"
fi
