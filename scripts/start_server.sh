#!/bin/bash

# Remove the previously files
rm app
sudo rm nohub.out

# Build the applicatiojn
go build -o ./app

# Start the server in the background
nohub ./app &
