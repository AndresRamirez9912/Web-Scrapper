#!/bin/sh

# Set execute permissions for the script
chmod +x "$0"

# Find the process ID (PID) of the server running on port 8080
PID=$(lsof -t -i :3000)

# Check if the server process is running
if [ -z "$PID" ]; then
  echo "Server is not running."
else
  # Stop the server gracefully by sending a termination signal
  kill "$PID"
  echo "Server stopped successfully."
fi
