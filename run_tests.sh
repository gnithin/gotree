#!/usr/bin/env bash

url="https://king-slayer.appspot.com/post"

cd ./tree/tests/
go test -v -bench . > output.out 2>&1

# Storing the exit code
exit_code="$?"

all_output=$(cat output.out)

# curl that resp
curl --data "allOutput=$all_output" "$url" > /dev/null 2>&1

# Exiting the code in the same way as the test resp
exit "$exit_code"
