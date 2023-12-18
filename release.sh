#!/usr/bin/env bash

# Colours
RED='\033[0;31m'
CYAN='\033[0;36m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

usage() {
	echo -e "${CYAN}usage: ./release.sh {1}${NC}"
	echo -e "${CYAN}parameters:${NC}"
	echo -e "  ${YELLOW}{1}     the version to release${NC}"
}

# check arguments
if [[ $# -ne 1 ]]; then
	usage
	exit 1
fi

# check tag doesn't already exist
if [[ ! -z $(git tag -l | grep $1) ]]; then
  echo -e "${RED}This version already exists${NC}"
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

ecrBase="294786226104.dkr.ecr.eu-west-2.amazonaws.com"

export AWS_PROFILE=ostop-prod

aws ecr get-login-password --region eu-west-2 | docker login --username AWS --password-stdin $ecrBase
docker build --platform=linux/amd64 -t $ecrBase/onestop-backend:$1 .
docker push $ecrBase/onestop-backend:$1
docker logout $ecrBase
docker rmi -f $(docker images "$ecrBase/onestop-backend" -a -q)

sed -E -i '' "s/api_service_version = \".+\"/api_service_version = \"$1\"/" _infrastructure/terraform/envs/prod/vars
sleep 3

docker build -f Dockerfile.release -t releaser .
docker run -v $HOME/.aws/onestop-config:/root/.aws/config -v $HOME/.aws/onestop-credentials:/root/.aws/credentials releaser
docker rmi -f $(docker images "releaser" -a -q)

git tag -a $1 -m \"$1\"
git push origin $1