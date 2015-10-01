#!/bin/bash

# Fail hard and fast
#set -eo pipefail

# Start healthchecker
echo "[healthchecker] starting healthchecker..."
/go/src/github.com/lukeatherton/healthchecker/healthchecker
