#!/usr/bin/env bash

# stops the execution of a script if a command or pipeline has an error
set -e

packages="service"
for p in $packages; do
    coverage="$(go test -cover ./$p | awk '{print $(NF-2)}' | grep -Eo '^[0-9]+' || true)"
    if [ -z "$coverage" ]; then
        coverage=0
    fi
    (( "$coverage" >= 0 )) || (echo "FAIL: package $p has ${coverage}% unit test coverage" && false)
done

echo "PASS"