#!/bin/sh

# Find the process ID (PID) of the server running on port 8080
PID=$(lsof -t -i :3000)

# Check if the server process is running
if ! [ -z "$PID" ]; then
  # Stop the server gracefully by sending a termination signal
  sudo kill "$PID"
  echo "Server stopped successfully."
fi
