#!/usr/bin/env bash

# Colours
RED='\033[0;31m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

usage() {
	echo -e "${CYAN}usage: ./infra.sh${NC}"
}

# check arguments
if [[ $# -ne 0 ]]; then
	usage
	exit 1
fi

# check aws config file exists
if [[ ! -f $HOME/.aws/onestop-config ]]; then
  echo -e "${RED}onestop-config file is required under the $HOME/.aws/ directory${NC}"
	exit 1
fi

# check aws credentials file exists
if [[ ! -f $HOME/.aws/onestop-credentials ]]; then
  echo -e "${RED}onestop-credentials file is required under the $HOME/.aws/ directory${NC}"
	exit 1
fi

docker build -f Dockerfile.infra -t infra .
docker run -v $HOME/.aws/onestop-config:/root/.aws/config -v $HOME/.aws/onestop-credentials:/root/.aws/credentials infra
docker rmi -f $(docker images "infra" -a -q)