#!/bin/bash

# Check if the binary file exists
if ! [ -f /home/ubuntu/price-tracking/Web-Scrapper/app ]; then
    echo "File does not exist"
    exit 1
fi

echo "File exists"

# Start the server in the background
nohup /home/ubuntu/price-tracking/Web-Scrapper/app > /dev/null 2>&1 & #Run in background but no generate output log 
exit 0
