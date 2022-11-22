#!/usr/bin/env bash
set -e

# e.g.,
# ./scripts/build.sh
# ./scripts/tests.upgrade.sh 1.7.16 ./build/coinflectchain
if ! [[ "$0" =~ scripts/tests.upgrade.sh ]]; then
  echo "must be run from repository root"
  exit 255
fi

VERSION=$1
if [[ -z "${VERSION}" ]]; then
  echo "Missing version argument!"
  echo "Usage: ${0} [VERSION] [NEW-BINARY]" >> /dev/stderr
  exit 255
fi

NEW_BINARY=$2
if [[ -z "${NEW_BINARY}" ]]; then
  echo "Missing new binary path argument!"
  echo "Usage: ${0} [VERSION] [NEW-BINARY]" >> /dev/stderr
  exit 255
fi

#################################
# download coinflectchain
# https://github.com/coinflect/coinflectchain/releases
GOARCH=$(go env GOARCH)
GOOS=$(go env GOOS)
DOWNLOAD_URL=https://github.com/coinflect/coinflectchain/releases/download/v${VERSION}/coinflectchain-linux-${GOARCH}-v${VERSION}.tar.gz
DOWNLOAD_PATH=/tmp/coinflectchain.tar.gz
if [[ ${GOOS} == "darwin" ]]; then
  DOWNLOAD_URL=https://github.com/coinflect/coinflectchain/releases/download/v${VERSION}/coinflectchain-macos-v${VERSION}.zip
  DOWNLOAD_PATH=/tmp/coinflectchain.zip
fi

rm -f ${DOWNLOAD_PATH}
rm -rf /tmp/coinflectchain-v${VERSION}
rm -rf /tmp/coinflectchain-build

echo "downloading coinflectchain ${VERSION} at ${DOWNLOAD_URL}"
curl -L ${DOWNLOAD_URL} -o ${DOWNLOAD_PATH}

echo "extracting downloaded coinflectchain"
if [[ ${GOOS} == "linux" ]]; then
  tar xzvf ${DOWNLOAD_PATH} -C /tmp
elif [[ ${GOOS} == "darwin" ]]; then
  unzip ${DOWNLOAD_PATH} -d /tmp/coinflectchain-build
  mv /tmp/coinflectchain-build/build /tmp/coinflectchain-v${VERSION}
fi
find /tmp/coinflectchain-v${VERSION}

#################################
# download coinflect-network-runner
# https://github.com/coinflect/coinflect-network-runner
NETWORK_RUNNER_VERSION=1.1.0
DOWNLOAD_PATH=/tmp/coinflect-network-runner.tar.gz
DOWNLOAD_URL=https://github.com/coinflect/coinflect-network-runner/releases/download/v${NETWORK_RUNNER_VERSION}/coinflect-network-runner_${NETWORK_RUNNER_VERSION}_linux_amd64.tar.gz
if [[ ${GOOS} == "darwin" ]]; then
  DOWNLOAD_URL=https://github.com/coinflect/coinflect-network-runner/releases/download/v${NETWORK_RUNNER_VERSION}/coinflect-network-runner_${NETWORK_RUNNER_VERSION}_darwin_amd64.tar.gz
fi

rm -f ${DOWNLOAD_PATH}
rm -f /tmp/coinflect-network-runner

echo "downloading coinflect-network-runner ${NETWORK_RUNNER_VERSION} at ${DOWNLOAD_URL}"
curl -L ${DOWNLOAD_URL} -o ${DOWNLOAD_PATH}

echo "extracting downloaded coinflect-network-runner"
tar xzvf ${DOWNLOAD_PATH} -C /tmp
/tmp/coinflect-network-runner -h

#################################
echo "building upgrade.test"
# to install the ginkgo binary (required for test build and run)
go install -v github.com/onsi/ginkgo/v2/ginkgo@v2.1.4
ACK_GINKGO_RC=true ginkgo build ./tests/upgrade
./tests/upgrade/upgrade.test --help

#################################
# run "coinflect-network-runner" server
echo "launch coinflect-network-runner in the background"
/tmp/coinflect-network-runner \
server \
--log-level debug \
--port=":12340" \
--disable-grpc-gateway &
PID=${!}

#################################
# By default, it runs all upgrade test cases!
echo "running upgrade tests against the local cluster with ${NEW_BINARY}"
./tests/upgrade/upgrade.test \
--ginkgo.v \
--log-level debug \
--network-runner-grpc-endpoint="0.0.0.0:12340" \
--network-runner-coinflectchain-path=/tmp/coinflectchain-v${VERSION}/coinflectchain \
--network-runner-coinflectchain-path-to-upgrade=${NEW_BINARY} \
--network-runner-coinflectchain-log-level="WARN" || EXIT_CODE=$?

# "e2e.test" already terminates the cluster
# just in case tests are aborted, manually terminate them again
pkill -P ${PID} || true
kill -2 ${PID}

if [[ ${EXIT_CODE} -gt 0 ]]; then
  echo "FAILURE with exit code ${EXIT_CODE}"
  exit ${EXIT_CODE}
else
  echo "ALL SUCCESS!"
fi
