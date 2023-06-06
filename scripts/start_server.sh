#!/bin/bash

# Remove the previously files
rm app
sudo rm nohup.out

# Build the applicatiojn
go build -o ./app

# Start the server in the background
nohup ./app &!
