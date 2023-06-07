#!/bin/bash

if [ -f /home/ubuntu/price-tracking/Web-Scrapper/app ]; then
    echo "File exists"
else
    echo "File does not exist"
    exit 1
fi

# Start the server in the background
/home/ubuntu/price-tracking/Web-Scrapper/app &
