#!/bin/bash

# Check if the binary file exists
if ! [ -f /home/ubuntu/price-tracking/Web-Scrapper/app ]; then
    echo "File does not exist"
    exit 1
fi

echo "File exists"

# Start the server in the background
/home/ubuntu/price-tracking/Web-Scrapper/app &
exit 0
