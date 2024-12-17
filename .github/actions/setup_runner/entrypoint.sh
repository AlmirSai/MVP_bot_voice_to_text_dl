#!/bin/bash
# Entrypoint script for the GitHub Runner setup action

# Check if the GitHub token and runner name are provided
if [[ -z "$1" || -z "$2" ]]; then
  echo "Runner token and name must be provided."
  exit 1
fi

# Setup environment variables
RUNNER_TOKEN=$1
RUNNER_NAME=$2
RUNNER_WORKDIR=${3:-/home/runner/work}

# Log the provided details
echo "Setting up GitHub Runner with token: $RUNNER_TOKEN, name: $RUNNER_NAME, workdir: $RUNNER_WORKDIR"

# Proceed with setting up the runner
./config.sh --url https://github.com --token $RUNNER_TOKEN --name $RUNNER_NAME --work $RUNNER_WORKDIR
./run.sh
