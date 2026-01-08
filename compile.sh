#!/bin/sh
#
# This script is used to compile your program on CodeCrafters
#
# This runs before .codecrafters/run.sh
#
# Learn more: https://codecrafters.io/program-interface

set -e # Exit on failure

(
  cd "$(dirname "$0")/.." # Change to repository root
  go build -o /tmp/codecrafters-build-shell-go app/*.go
)
