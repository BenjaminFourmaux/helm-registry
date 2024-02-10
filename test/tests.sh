#!/bin/bash

# ===========
# Variables
# ===========
REGISTRY_URL="http://registry:8080"

# ===========
# Functions
# ===========

printTest() {
    echo "Running test: [$1] - $2"
}


# ===========
# Tests
# ===========

# Test 0 - Check API is up
printTest "0" "Check API is up"
if ! curl -s -o /dev/null -w "%{http_code}" $REGISTRY_URL | grep -q "200"; then
    echo "Registry is down"
    exit 1
else
    echo "Registry is up"
fi

# Test 1 - Add repository
printTest "1" "Add repository"
if helm repo add test $REGISTRY_URL; then
    echo "Repository added"
else
    echo "Repository not added"
    exit 1
fi

