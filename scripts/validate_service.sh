#!/bin/bash

# Check if the server is running on port 3000
if ! nc -z localhost 3000; then
  echo "Service is not running."
  exit 1  # Return a non-zero exit code to indicate failure
fi
