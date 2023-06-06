#!/bin/bash

# Go to the application path
cd /home/ubuntu/price-tracking/Web-Scrapper/

# Remove the previously files
rm app

# Build the applicatiojn
go build -o ./app
if [ -f ./app ]; then
    echo "File exists"
else
    echo "File does not exist"
    exit 1
fi

# Start the server in the background
./app &
